package cart

import (
	"errors"

	"github.com/havilz/ecommerce-backend/internal/modules/product"
	"github.com/havilz/ecommerce-backend/models"
	"gorm.io/gorm"
)

type Service interface {
	GetCart(userID uint) (*CartResponse, error)
	AddItem(userID uint, req AddCartRequest) error
	UpdateItem(userID, itemID uint, req UpdateCartRequest) error
	RemoveItem(userID, itemID uint) error
	ClearCart(userID uint) error
}

type service struct {
	repo        Repository
	productRepo product.Repository
}

func NewService(repo Repository, productRepo product.Repository) Service {
	return &service{repo: repo, productRepo: productRepo}
}

func (s *service) GetCart(userID uint) (*CartResponse, error) {
	items, err := s.repo.FindByUserID(userID)
	if err != nil {
		return nil, err
	}

	var responseItems []CartItemResponse
	for _, item := range items {
		responseItems = append(responseItems, CartItemResponse{
			ID: item.ID,
			Product: Product{
				ID:       item.Product.ID,
				Name:     item.Product.Name,
				Price:    item.Product.Price,
				Stock:    item.Product.Stock,
				ImageURL: item.Product.ImageURL,
				Category: item.Product.Category,
			},
			Quantity: item.Quantity,
		})
	}

	if responseItems == nil {
		responseItems = []CartItemResponse{}
	}

	return &CartResponse{
		Items:      responseItems,
		TotalItems: len(responseItems),
	}, nil
}

func (s *service) AddItem(userID uint, req AddCartRequest) error {
	if req.Quantity < 1 {
		return errors.New("quantity must be at least 1")
	}

	prod, err := s.productRepo.FindByID(req.ProductID)
	if err != nil {
		return errors.New("product not found")
	}

	// If cart item already exists, merge quantity
	existing, err := s.repo.FindByUserAndProduct(userID, req.ProductID)
	if err == nil {
		newQty := existing.Quantity + req.Quantity
		if prod.Stock < newQty {
			return errors.New("insufficient stock")
		}
		existing.Quantity = newQty
		return s.repo.Update(existing)
	}

	if !errors.Is(err, gorm.ErrRecordNotFound) {
		return err
	}

	if prod.Stock < req.Quantity {
		return errors.New("insufficient stock")
	}

	return s.repo.Create(&models.CartItem{
		UserID:    userID,
		ProductID: req.ProductID,
		Quantity:  req.Quantity,
	})
}

func (s *service) UpdateItem(userID, itemID uint, req UpdateCartRequest) error {
	if req.Quantity < 1 {
		return errors.New("quantity must be at least 1")
	}

	item, err := s.repo.FindItemByID(itemID, userID)
	if err != nil {
		return errors.New("cart item not found")
	}

	prod, err := s.productRepo.FindByID(item.ProductID)
	if err != nil {
		return errors.New("product not found")
	}
	if prod.Stock < req.Quantity {
		return errors.New("insufficient stock")
	}

	item.Quantity = req.Quantity
	return s.repo.Update(item)
}

func (s *service) RemoveItem(userID, itemID uint) error {
	return s.repo.Delete(itemID, userID)
}

func (s *service) ClearCart(userID uint) error {
	return s.repo.ClearByUserID(userID)
}
