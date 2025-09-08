package cartRepository

import (
	"github.com/meshyampratap01/OnlineShoppingCart/internal/dto"
	"github.com/meshyampratap01/OnlineShoppingCart/internal/models"
)

type CartManager interface {
	CreateCart(cartID, userID string) error
	GetCartIDByUserID(userID string) (string, error)
	AddToCart(userID string, product models.Product) error
	RemoveFromCart(cartID string, prodID string) error
	EmptyCart(userID string) error
	GetCartItemQuantity(cartID,prodID string) (int,error)
	GetCartItems(cartID string) ([]dto.CartItemsDTO, error)
}
