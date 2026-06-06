package models

func AllModels() []interface{} {
	return []interface{}{
		&User{},
		&Product{},
		&CartItem{},
		&Order{},
		&OrderItem{},
	}
}