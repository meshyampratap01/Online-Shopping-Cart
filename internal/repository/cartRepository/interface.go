package cartrepository

import (
	"github.com/meshyampratap01/OnlineShoppingCart/internal/dto"
	"github.com/meshyampratap01/OnlineShoppingCart/internal/models"
)

type CartManager interface {
	GetCartIDByUserID(userID string) (string, error)
	AddToCart(userID string, product models.Product) error
	RemoveFromCart(userID string, prodID string) error
	EmptyCart(userID string) error
	GetCartItems(userID string) ([]dto.CartItemsDTO, error)
}
