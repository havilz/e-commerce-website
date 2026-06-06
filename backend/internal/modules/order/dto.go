package order

import "time"

type CheckoutRequest struct {
	Address string `json:"address"`
}

type OrderProduct struct {
	ID       uint    `json:"id"`
	Name     string  `json:"name"`
	Price    float64 `json:"price"`
	ImageURL string  `json:"image_url"`
	Category string  `json:"category"`
}

type OrderItemResponse struct {
	ID              uint         `json:"id"`
	Product         OrderProduct `json:"product"`
	Quantity        int          `json:"quantity"`
	PriceAtPurchase float64      `json:"price_at_purchase"`
}

type OrderResponse struct {
	ID         uint                `json:"id"`
	TotalPrice float64             `json:"total_price"`
	Status     string              `json:"status"`
	Address    string              `json:"address"`
	Items      []OrderItemResponse `json:"items"`
	CreatedAt  time.Time           `json:"created_at"`
}
