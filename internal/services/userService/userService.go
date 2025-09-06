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
	cart := []models.Product{}
	newUser := models.User{
		ID:       id,
		Name:     name,
		Email:    email,
		Password: password,
		Cart:     cart,
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

func (us *UserService) AddToCart(userID, prodID string) error {
	prod, err := us.prodRepo.GetProductByID(prodID)
	if err != nil {
		return fmt.Errorf("no product with specified id found")
	}
	err = us.userRepo.AddToUserCart(userID, prod)
	if err != nil {
		return fmt.Errorf("can't add product to cart")
	}
	return nil
}

func (us *UserService) RemoveFromCart(userID, prodID string) error {
	prod, err := us.prodRepo.GetProductByID(prodID)
	if err != nil {
		return fmt.Errorf("no product with specified id found")
	}
	cart, err := us.userRepo.GetUserCart(userID)
	if err != nil {
		return fmt.Errorf("can not fetch cart for user")
	}
	if len(cart) == 0 {
		return fmt.Errorf("cart is empty")
	}
	err = us.userRepo.RemoveFromUserCart(userID, prod)
	if err != nil {
		return fmt.Errorf("can't remove product from cart")
	}
	return nil
}

func (us *UserService) GetCartByUserID(id string) ([]models.Product, float32, error) {
	cart, err := us.userRepo.GetUserCart(id)
	if err != nil {
		return nil, 0, fmt.Errorf("can not fetch cart for user")
	}
	var totalPrice float32
	for _, prod := range cart {
		totalPrice += prod.Price
	}
	return cart, totalPrice, nil
}

func (us *UserService) CheckOut(id string, couponCode string) (float32, error) {
	cart, err := us.userRepo.GetUserCart(id)
	if err != nil {
		return 0, fmt.Errorf("can not fetch cart for user")
	}
	if len(cart) == 0 {
		return 0, fmt.Errorf("cart is empty")
	}
	var total float32
	for _, prod := range cart {
		if prod.Stock <= 0 {
			return 0, fmt.Errorf("product %s is out of stock", prod.Name)
		} else {
			prod.Stock -= 1
			err := us.prodRepo.UpdateProduct(prod)
			if err != nil {
				return 0, fmt.Errorf("can not update product %s", prod.Name)
			}
		}
		total += prod.Price
	}
	if couponCode != "" {
		coupon, err := us.couponRepo.GetCouponByCode(couponCode)
		if err != nil || coupon == nil {
			return 0, fmt.Errorf("invalid coupon code")
		}
		total *= (1 - coupon.Percentage/100)
	}

	return total, nil
}
