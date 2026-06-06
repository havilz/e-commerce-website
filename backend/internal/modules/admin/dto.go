package admin

import "time"

type AdminProductRequest struct {
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Price       float64 `json:"price"`
	Stock       int     `json:"stock"`
	ImageURL    string  `json:"image_url"`
	Category    string  `json:"category"`
}

type UpdateOrderStatusRequest struct {
	Status string `json:"status"`
}

type AdminProductResponse struct {
	ID          uint    `json:"id"`
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Price       float64 `json:"price"`
	Stock       int     `json:"stock"`
	ImageURL    string  `json:"image_url"`
	Category    string  `json:"category"`
}

type AdminOrderProduct struct {
	ID       uint    `json:"id"`
	Name     string  `json:"name"`
	Price    float64 `json:"price"`
	ImageURL string  `json:"image_url"`
}

type AdminOrderItemResponse struct {
	ID              uint              `json:"id"`
	Product         AdminOrderProduct `json:"product"`
	Quantity        int               `json:"quantity"`
	PriceAtPurchase float64           `json:"price_at_purchase"`
}

type AdminOrderUser struct {
	ID    uint   `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

type AdminOrderResponse struct {
	ID         uint                     `json:"id"`
	UserID     uint                     `json:"user_id"`
	User       AdminOrderUser           `json:"user"`
	TotalPrice float64                  `json:"total_price"`
	Status     string                   `json:"status"`
	Address    string                   `json:"address"`
	Items      []AdminOrderItemResponse  `json:"items"`
	CreatedAt  time.Time                `json:"created_at"`
}
