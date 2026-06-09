package models

func AllModels() []interface{} {
	return []interface{}{
		&User{},
		&Category{},
		&Product{},
		&CartItem{},
		&Order{},
		&OrderItem{},
	}
}