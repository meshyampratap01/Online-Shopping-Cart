package models

type Product struct {
	ID    string  `json:"id"`
	Name  string  `json:"name"`
	Price float32 `json:"price"`
	Stock int     `json:"stock"`
}
