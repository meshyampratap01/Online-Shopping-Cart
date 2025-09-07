package userService

import (
	"fmt"

	"github.com/meshyampratap01/OnlineShoppingCart/internal/models"
	"github.com/meshyampratap01/OnlineShoppingCart/internal/repository/couponRepository"
	"github.com/meshyampratap01/OnlineShoppingCart/internal/repository/productRepository"
	"github.com/meshyampratap01/OnlineShoppingCart/internal/repository/userRepository"
	"github.com/meshyampratap01/OnlineShoppingCart/internal/utils"
)

type UserService struct {
	userRepo   userRepository.UserManager
	prodRepo   productRepository.ProductManager
	couponRepo couponRepository.CouponManager
}

func NewUserService(userRepo userRepository.UserManager, prodRepo productRepository.ProductManager, couponRepo couponRepository.CouponManager) UserServiceManager {
	return &UserService{
		userRepo:   userRepo,
		prodRepo:   prodRepo,
		couponRepo: couponRepo,
	}
}

func (us *UserService) RegisterUser(name, email, password string, role models.UserRole) error {
	user, err := us.userRepo.GetUserByEmail(email)
	if err == nil && user.Email == email {
		return fmt.Errorf("user with email %s already exists", email)
	}

	newUser, err := us.CreateUser(name, email, password, role)
	if err != nil {
		return fmt.Errorf("can not create new user")
	}

	err = us.userRepo.SaveUser(newUser)
	if err != nil {
		return fmt.Errorf("can not save new user")
	}

	return nil
}

func (us *UserService) CreateUser(name, email, password string, role models.UserRole) (models.User, error) {
	id := utils.NewUUID()
	newUser := models.User{
		ID:       id,
		Name:     name,
		Email:    email,
		Password: password,
	}

	return newUser, nil
}

func (us *UserService) Login(email, password string) (string, error) {
	user, err := us.userRepo.GetUserByEmail(email)
	if err != nil {
		return "", fmt.Errorf("invalid email or password")
	}
	if !utils.CheckPassword(user.Password, password) {
		return "", fmt.Errorf("invalid email or password")
	}

	userJWT := models.UserJWT{
		UserID: user.ID,
		Email:  user.Email,
		Role:   user.Role,
	}
	token, err := utils.GenerateJWT(userJWT)
	if err != nil {
		return "nil", fmt.Errorf("can not generate token")
	}

	return token, nil
}