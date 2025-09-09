package cartservice

import (
	"github.com/meshyampratap01/OnlineShoppingCart/internal/dto"
)

//go:generate mockgen -source=interface.go -destination=../../mocks/mock_cartServcie.go -package mocks

type CartServiceManager interface {
	GetCartItems(userID string) ([]dto.CartItemsDTO, error)
	AddToCart(userID, prodID string) error
	RemoveFromCart(userID, prodID string) error
	Checkout(userID string, couponCode string) (float32, error)
}