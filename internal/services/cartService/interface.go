package cartservice

import (
	"github.com/meshyampratap01/OnlineShoppingCart/internal/dto"
)

type CartServiceManager interface {
	GetCartItems(cartID string) ([]dto.CartItemsDTO, error)
	AddToCart(userID, prodID string) error
	RemoveFromCart(userID, prodID string) error
	Checkout(userID string, couponCode string) (float32, error)
}