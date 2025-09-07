package userRepository

import "github.com/meshyampratap01/OnlineShoppingCart/internal/models"

type UserManager interface {
	SaveUser(models.User) error
	GetUserByID(id string) (models.User, error)
	GetUserByEmail(email string) (models.User, error)
}
