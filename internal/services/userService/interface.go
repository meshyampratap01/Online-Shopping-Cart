package userService

import "github.com/meshyampratap01/OnlineShoppingCart/internal/models"

type UserServiceManager interface {
	RegisterUser(name, email, password string, role models.UserRole) error
	Login(email, password string) (string, error)
}
