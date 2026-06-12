package admin

import (
	"errors"
	"strings"

	"github.com/havilz/ecommerce-backend/models"
	"github.com/havilz/ecommerce-backend/pkg/security"
)

var validStatuses = map[string]bool{
	"pending":    true,
	"processing": true,
	"completed":  true,
	"cancelled":  true,
}

type Service interface {
	CreateProduct(req AdminProductRequest) (*AdminProductResponse, error)
	UpdateProduct(id uint, req AdminProductRequest) error
	DeleteProduct(id uint) error
	GetAllOrders(page, limit int) ([]AdminOrderResponse, int64, error)
	UpdateOrderStatus(orderID uint, req UpdateOrderStatusRequest) error
}

type service struct {
	repo Repository
}

func NewService(repo Repository) Service {
	return &service{repo: repo}
}

func (s *service) CreateProduct(req AdminProductRequest) (*AdminProductResponse, error) {
	// Sanitize text fields
	req.Name = security.SanitizeString(req.Name)
	req.Description = security.SanitizeString(req.Description)
	req.ImageURL = strings.TrimSpace(req.ImageURL)

	if req.Name == "" || req.Price <= 0 {
		return nil, errors.New("name and a positive price are required")
	}

	// Validate ImageURL scheme if provided to prevent URI-based XSS (e.g. javascript:...)
	if req.ImageURL != "" {
		lowerURL := strings.ToLower(req.ImageURL)
		if !strings.HasPrefix(lowerURL, "http://") && !strings.HasPrefix(lowerURL, "https://") {
			return nil, errors.New("image_url must start with http:// or https://")
		}
	}

	var categoryID uint
	var catName string
	if req.CategoryID > 0 {
		cat, err := s.repo.FindCategoryByID(req.CategoryID)
		if err != nil {
			return nil, errors.New("category not found")
		}
		categoryID = cat.ID
		catName = cat.Name
	} else if req.Category != "" {
		cat, err := s.repo.FindCategoryByName(req.Category)
		if err != nil {
			return nil, errors.New("category '" + req.Category + "' not found. Admin must create this category first")
		}
		categoryID = cat.ID
		catName = cat.Name
	} else {
		return nil, errors.New("category is required")
	}

	p := &models.Product{
		Name:        req.Name,
		Description: req.Description,
		Price:       req.Price,
		Stock:       req.Stock,
		ImageURL:    req.ImageURL,
		CategoryID:  categoryID,
	}

	if err := s.repo.CreateProduct(p); err != nil {
		return nil, errors.New("failed to create product")
	}

	return &AdminProductResponse{
		ID:          p.ID,
		Name:        p.Name,
		Description: p.Description,
		Price:       p.Price,
		Stock:       p.Stock,
		ImageURL:    p.ImageURL,
		CategoryID:  p.CategoryID,
		Category:    catName,
	}, nil
}

func (s *service) UpdateProduct(id uint, req AdminProductRequest) error {
	// Sanitize text fields
	req.Name = security.SanitizeString(req.Name)
	req.Description = security.SanitizeString(req.Description)
	req.ImageURL = strings.TrimSpace(req.ImageURL)

	if req.Name == "" || req.Price <= 0 {
		return errors.New("name and a positive price are required")
	}

	// Validate ImageURL scheme if provided to prevent URI-based XSS (e.g. javascript:...)
	if req.ImageURL != "" {
		lowerURL := strings.ToLower(req.ImageURL)
		if !strings.HasPrefix(lowerURL, "http://") && !strings.HasPrefix(lowerURL, "https://") {
			return errors.New("image_url must start with http:// or https://")
		}
	}

	var categoryID uint
	if req.CategoryID > 0 {
		cat, err := s.repo.FindCategoryByID(req.CategoryID)
		if err != nil {
			return errors.New("category not found")
		}
		categoryID = cat.ID
	} else if req.Category != "" {
		cat, err := s.repo.FindCategoryByName(req.Category)
		if err != nil {
			return errors.New("category '" + req.Category + "' not found. Admin must create this category first")
		}
		categoryID = cat.ID
	} else {
		return errors.New("category is required")
	}

	data := map[string]interface{}{
		"name":        req.Name,
		"description": req.Description,
		"price":       req.Price,
		"stock":       req.Stock,
		"image_url":   req.ImageURL,
		"category_id": categoryID,
	}
	return s.repo.UpdateProduct(id, data)
}

func (s *service) DeleteProduct(id uint) error {
	return s.repo.DeleteProduct(id)
}

func (s *service) GetAllOrders(page, limit int) ([]AdminOrderResponse, int64, error) {
	if page < 1 {
		page = 1
	}
	if limit < 1 || limit > 100 {
		limit = 10
	}

	orders, total, err := s.repo.FindAllOrders(page, limit)
	if err != nil {
		return nil, 0, err
	}

	var result []AdminOrderResponse
	for _, o := range orders {
		var items []AdminOrderItemResponse
		for _, item := range o.Items {
			items = append(items, AdminOrderItemResponse{
				ID: item.ID,
				Product: AdminOrderProduct{
					ID:       item.Product.ID,
					Name:     item.Product.Name,
					Price:    item.Product.Price,
					ImageURL: item.Product.ImageURL,
				},
				Quantity:        item.Quantity,
				PriceAtPurchase: item.PriceAtPurchase,
			})
		}
		if items == nil {
			items = []AdminOrderItemResponse{}
		}
		result = append(result, AdminOrderResponse{
			ID:     o.ID,
			UserID: o.UserID,
			User: AdminOrderUser{
				ID:    o.User.ID,
				Name:  o.User.Name,
				Email: o.User.Email,
			},
			TotalPrice: o.TotalPrice,
			Status:     o.Status,
			Address:    o.Address,
			Items:      items,
			CreatedAt:  o.CreatedAt,
		})
	}
	if result == nil {
		result = []AdminOrderResponse{}
	}
	return result, total, nil
}

func (s *service) UpdateOrderStatus(orderID uint, req UpdateOrderStatusRequest) error {
	if !validStatuses[req.Status] {
		return errors.New("invalid status: must be pending, processing, completed, or cancelled")
	}
	return s.repo.UpdateOrderStatus(orderID, req.Status)
}
