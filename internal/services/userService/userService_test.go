package userService

import (
    "errors"
    "testing"

    "github.com/meshyampratap01/OnlineShoppingCart/internal/mocks"
    "github.com/meshyampratap01/OnlineShoppingCart/internal/models"
    "github.com/meshyampratap01/OnlineShoppingCart/internal/utils"
    "go.uber.org/mock/gomock"
)

func TestRegisterUser(t *testing.T) {
    ctrl := gomock.NewController(t)
    defer ctrl.Finish()

    mockUserRepo := mocks.NewMockUserManager(ctrl)
    mockProdRepo := mocks.NewMockProductManager(ctrl)
    mockCouponRepo := mocks.NewMockCouponManager(ctrl)
    mockCartRepo := mocks.NewMockCartManager(ctrl)

    service := NewUserService(mockUserRepo, mockProdRepo, mockCouponRepo, mockCartRepo)

    email := "test@example.com"
    name := "Test User"
    password := "password123"
    role := models.Customer

    t.Run("User already exists", func(t *testing.T) {
        mockUserRepo.EXPECT().GetUserByEmail(email).Return(models.User{Email: email}, nil)

        err := service.RegisterUser(name, email, password, role)
        if err == nil {
            t.Errorf("expected error for existing user, got nil")
        }
    })

    t.Run("Error saving user", func(t *testing.T) {
        mockUserRepo.EXPECT().GetUserByEmail(email).Return(models.User{}, errors.New("not found"))
        mockUserRepo.EXPECT().SaveUser(gomock.Any()).Return(errors.New("save error"))

        err := service.RegisterUser(name, email, password, role)
        if err == nil {
            t.Errorf("expected error during saving user, got nil")
        }
    })

    t.Run("Error creating cart", func(t *testing.T) {
        mockUserRepo.EXPECT().GetUserByEmail(email).Return(models.User{}, errors.New("not found"))
        mockUserRepo.EXPECT().SaveUser(gomock.Any()).Return(nil)
        mockCartRepo.EXPECT().CreateCart(gomock.Any(), gomock.Any()).Return(errors.New("cart error"))

        err := service.RegisterUser(name, email, password, role)
        if err == nil {
            t.Errorf("expected error during cart creation, got nil")
        }
    })

    t.Run("Successful registration", func(t *testing.T) {
        mockUserRepo.EXPECT().GetUserByEmail(email).Return(models.User{}, errors.New("not found"))
        mockUserRepo.EXPECT().SaveUser(gomock.Any()).Return(nil)
        mockCartRepo.EXPECT().CreateCart(gomock.Any(), gomock.Any()).Return(nil)

        err := service.RegisterUser(name, email, password, role)
        if err != nil {
            t.Errorf("expected successful registration, got error: %v", err)
        }
    })
}

func TestCreateUser(t *testing.T) {
    service := UserService{}

    t.Run("Success", func(t *testing.T) {
        user, err := service.CreateUser("name", "email", "pass", models.Customer)
        if err != nil {
            t.Errorf("expected success, got error: %v", err)
        }
        if user.Email != "email" {
            t.Errorf("expected email to be 'email', got %s", user.Email)
        }
    })
}

func TestLogin(t *testing.T) {
    ctrl := gomock.NewController(t)
    defer ctrl.Finish()

    mockUserRepo := mocks.NewMockUserManager(ctrl)
    service := UserService{userRepo: mockUserRepo}

    email := "test@example.com"
    password := "password123"
    hashedPassword, _ := utils.HashPassword(password)

    t.Run("Invalid email", func(t *testing.T) {
        mockUserRepo.EXPECT().GetUserByEmail(email).Return(models.User{}, errors.New("not found"))

        _, err := service.Login(email, password)
        if err == nil {
            t.Errorf("expected error for invalid email, got nil")
        }
    })

    t.Run("Invalid password", func(t *testing.T) {
        mockUserRepo.EXPECT().GetUserByEmail(email).Return(models.User{Password: "wronghash"}, nil)

        _, err := service.Login(email, password)
        if err == nil {
            t.Errorf("expected error for invalid password, got nil")
        }
    })

    t.Run("Successful login", func(t *testing.T) {
        mockUserRepo.EXPECT().GetUserByEmail(email).Return(models.User{
            ID:       "1",
            Email:    email,
            Password: hashedPassword,
            Role:     models.Customer,
        }, nil)

        token, err := service.Login(email, password)
        if err != nil {
            t.Errorf("expected successful login, got error: %v", err)
        }
        if token == "" {
            t.Errorf("expected non-empty token, got empty string")
        }
    })
}
