package order

import (
	"errors"

	cartRepo "github.com/havilz/ecommerce-backend/internal/modules/cart"
	productRepo "github.com/havilz/ecommerce-backend/internal/modules/product"
	"github.com/havilz/ecommerce-backend/models"
	"github.com/havilz/ecommerce-backend/pkg/security"
	"gorm.io/gorm"
)

type Service interface {
	Checkout(userID uint, req CheckoutRequest) (*OrderResponse, error)
	GetOrders(userID uint) ([]OrderResponse, error)
	GetOrderByID(orderID, userID uint) (*OrderResponse, error)
}

type service struct {
	repo        Repository
	cartRepo    cartRepo.Repository
	productRepo productRepo.Repository
	db          *gorm.DB
}

func NewService(
	repo Repository,
	cart cartRepo.Repository,
	product productRepo.Repository,
	db *gorm.DB,
) Service {
	return &service{
		repo:        repo,
		cartRepo:    cart,
		productRepo: product,
		db:          db,
	}
}

func (s *service) Checkout(userID uint, req CheckoutRequest) (*OrderResponse, error) {
	req.Address = security.SanitizeString(req.Address)
	if req.Address == "" {
		return nil, errors.New("address is required")
	}

	cartItems, err := s.cartRepo.FindByUserID(userID)
	if err != nil || len(cartItems) == 0 {
		return nil, errors.New("cart is empty")
	}

	var createdOrder *models.Order

	txErr := s.db.Transaction(func(tx *gorm.DB) error {
		var totalPrice float64
		var orderItems []models.OrderItem

		for _, item := range cartItems {
			prod, err := s.productRepo.FindByID(item.ProductID)
			if err != nil {
				return errors.New("product not found")
			}
			if prod.Stock < item.Quantity {
				return errors.New("insufficient stock for: " + prod.Name)
			}

			// Deduct stock
			if err := tx.Model(&models.Product{}).
				Where("id = ?", prod.ID).
				UpdateColumn("stock", gorm.Expr("stock - ?", item.Quantity)).Error; err != nil {
				return err
			}

			totalPrice += prod.Price * float64(item.Quantity)

			orderItems = append(orderItems, models.OrderItem{
				ProductID:       prod.ID,
				Quantity:        item.Quantity,
				PriceAtPurchase: prod.Price,
			})
		}

		order := &models.Order{
			UserID:     userID,
			TotalPrice: totalPrice,
			Status:     "pending",
			Address:    req.Address,
			Items:      orderItems,
		}

		if err := tx.Create(order).Error; err != nil {
			return err
		}

		// Clear cart inside transaction
		if err := tx.Where("user_id = ?", userID).Delete(&models.CartItem{}).Error; err != nil {
			return err
		}

		createdOrder = order
		return nil
	})

	if txErr != nil {
		return nil, txErr
	}

	// Re-fetch with preloaded items for the response
	fullOrder, err := s.repo.FindByID(createdOrder.ID, userID)
	if err != nil {
		return toOrderResponse(createdOrder), nil
	}

	return toOrderResponse(fullOrder), nil
}

func (s *service) GetOrders(userID uint) ([]OrderResponse, error) {
	orders, err := s.repo.FindByUserID(userID)
	if err != nil {
		return nil, err
	}

	var result []OrderResponse
	for _, o := range orders {
		result = append(result, *toOrderResponse(&o))
	}
	if result == nil {
		result = []OrderResponse{}
	}
	return result, nil
}

func (s *service) GetOrderByID(orderID, userID uint) (*OrderResponse, error) {
	order, err := s.repo.FindByID(orderID, userID)
	if err != nil {
		return nil, errors.New("order not found")
	}
	return toOrderResponse(order), nil
}

func toOrderResponse(o *models.Order) *OrderResponse {
	var items []OrderItemResponse
	for _, item := range o.Items {
		items = append(items, OrderItemResponse{
			ID: item.ID,
			Product: OrderProduct{
				ID:       item.Product.ID,
				Name:     item.Product.Name,
				Price:    item.Product.Price,
				ImageURL: item.Product.ImageURL,
				Category: item.Product.Category.Name,
			},
			Quantity:        item.Quantity,
			PriceAtPurchase: item.PriceAtPurchase,
		})
	}
	if items == nil {
		items = []OrderItemResponse{}
	}
	return &OrderResponse{
		ID:         o.ID,
		TotalPrice: o.TotalPrice,
		Status:     o.Status,
		Address:    o.Address,
		Items:      items,
		CreatedAt:  o.CreatedAt,
	}
}
