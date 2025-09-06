package userService

import "github.com/meshyampratap01/OnlineShoppingCart/internal/models"

type UserServiceManager interface {
	RegisterUser(name, email, password string, role models.UserRole) error
	Login(email, password string) (string, error)
	AddToCart(userID, prodID string) error
	RemoveFromCart(userID, prodID string) error
	GetCartByUserID(id string) ([]models.Product, float32, error)
	CheckOut(id string, couponCode string) (float32, error)
}
