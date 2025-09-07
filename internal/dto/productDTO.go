package dto

type ProductDTO struct {
	Name  string  `json:"name,omitempty"`
	Price float32 `json:"price,omitempty"`
	Stock int     `json:"stock,omitempty"`
}
