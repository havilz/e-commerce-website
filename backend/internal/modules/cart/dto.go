package cart

type AddCartRequest struct {
	ProductID uint `json:"product_id"`
	Quantity  int  `json:"quantity"`
}

type UpdateCartRequest struct {
	Quantity int `json:"quantity"`
}

type CartItemResponse struct {
	ID       uint    `json:"id"`
	Product  Product `json:"product"`
	Quantity int     `json:"quantity"`
}

// Product is an embedded product shape used inside cart responses
// to avoid circular imports with the product module.
type Product struct {
	ID       uint    `json:"id"`
	Name     string  `json:"name"`
	Price    float64 `json:"price"`
	Stock    int     `json:"stock"`
	ImageURL string  `json:"image_url"`
	Category string  `json:"category"`
}

type CartResponse struct {
	Items      []CartItemResponse `json:"items"`
	TotalItems int                `json:"total_items"`
}
