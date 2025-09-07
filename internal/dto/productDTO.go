package dto

type ProductDTO struct {
	Name  string  `json:"name"`
	Price float32 `json:"price"`
	Stock int     `json:"stock"`
}
