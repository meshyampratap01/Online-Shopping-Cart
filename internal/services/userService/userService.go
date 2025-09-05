package userService

import (
	"fmt"

	"github.com/meshyampratap01/OnlineShoppingCart/internal/models"
	"github.com/meshyampratap01/OnlineShoppingCart/internal/repository/userRepository"
	"github.com/meshyampratap01/OnlineShoppingCart/internal/utils"
)

type UserService struct {
	userRepo userRepository.UserManager
}

func NewUserService(userRepo userRepository.UserManager) *UserService {
	return &UserService{
		userRepo: userRepo,
	}
}

func (us *UserService) Signup(name, email, password string, role models.UserRole) error {
	//Todo validate email and password

	_, err := us.CreateUser(name, email, password, role)
	if err != nil {
		return fmt.Errorf("can not create new user")
	}

	//Todo save the user with the help of repo
	return nil
}

func (us *UserService) CreateUser(name, email, password string, role models.UserRole) (*models.User, error) {
	id, err := utils.NewUUID()
	if err != nil {
		return nil, fmt.Errorf("can't generate uuid")
	}
	cart := []models.Product{}
	newUser := &models.User{
		ID:       id,
		Name:     name,
		Email:    email,
		Password: password,
		Cart:     cart,
	}

	return newUser, nil
}

func (us *UserService) Login(email, password string) (*models.User, error) {
	//todo: fetch user from the database and then authenticate him/her
	return nil, nil
}
