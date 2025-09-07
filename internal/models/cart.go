package models

type Cart struct {
	ID     string `json:"id"`
	UserID string `json:"user_id"`
}

type CartItems struct {
	ID        string `json:"id"`
	CartID    string `json:"cart_id"`
	ProductID string `json:"product_id"`
	Quantity  int    `json:"quantity"`
}
