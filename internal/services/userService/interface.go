package userService

import "github.com/meshyampratap01/OnlineShoppingCart/internal/models"

//go:generate mockgen -source=interface.go -destination=../../mocks/mock_userServcie.go -package mocks

type UserServiceManager interface {
	RegisterUser(name, email, password string, role models.UserRole) error
	Login(email, password string) (string, error)
}
