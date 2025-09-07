package dto

type CartItemsDTO struct {
	ProductID   string  `json:"product_id"`
	ProductName string  `json:"product_name"`
	Price      float32 `json:"price"`
	Quantity   int     `json:"quantity"`
}