package userService

import "github.com/meshyampratap01/OnlineShoppingCart/internal/models"

type UserServiceManager interface {
	Signup(name, email, password string, role models.UserRole) error
	Login(email, password string) (*models.User, error)
	AddToCart(id string) error
	RemoveFromCart(id string) error
	GetCartByUserID(id string) ([]models.Product, error)
	CheckOut(id string) (int, error)
}
