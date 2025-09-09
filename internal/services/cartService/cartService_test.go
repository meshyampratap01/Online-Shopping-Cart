package cartservice

import (
	"errors"
	"testing"

	"github.com/meshyampratap01/OnlineShoppingCart/internal/dto"
	"github.com/meshyampratap01/OnlineShoppingCart/internal/mocks"
	"github.com/meshyampratap01/OnlineShoppingCart/internal/models"
	"go.uber.org/mock/gomock"
)

func TestGetCartItems(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockCartRepo := mocks.NewMockCartManager(ctrl)
	mockProdRepo := mocks.NewMockProductManager(ctrl)
	mockCouponRepo := mocks.NewMockCouponManager(ctrl)
	service := NewCartService(mockCartRepo, mockProdRepo, mockCouponRepo)

	mockCartRepo.EXPECT().GetCartIDByUserID("user1").Return("cart123", nil)
	mockCartRepo.EXPECT().GetCartItems("cart123").Return([]dto.CartItemsDTO{
		{ProductID: "p1", ProductName: "Item1", Price: 100, Quantity: 2},
	}, nil)

	items, err := service.GetCartItems("user1")
	if err != nil || len(items) != 1 {
		t.Errorf("unexpected error or wrong item count: %v", err)
	}

	mockCartRepo.EXPECT().GetCartIDByUserID("user2").Return("", errors.New("not found"))
	_, err = service.GetCartItems("user2")
	if err == nil {
		t.Error("expected error for missing cart")
	}
}

func TestAddToCart(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockCartRepo := mocks.NewMockCartManager(ctrl)
	mockProdRepo := mocks.NewMockProductManager(ctrl)
	mockCouponRepo := mocks.NewMockCouponManager(ctrl)
	service := NewCartService(mockCartRepo, mockProdRepo, mockCouponRepo)

	product := models.Product{ID: "p1", Name: "Item1", Stock: 5}
	mockProdRepo.EXPECT().GetProductByID("p1").Return(product, nil)
	mockCartRepo.EXPECT().GetCartIDByUserID("user1").Return("cart123", nil)
	mockCartRepo.EXPECT().GetCartItemQuantity("cart123", "p1").Return(2, nil)
	mockCartRepo.EXPECT().AddToCart("user1", product).Return(nil)

	err := service.AddToCart("user1", "p1")
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}

	product.Stock = 0
	mockProdRepo.EXPECT().GetProductByID("p2").Return(product, nil)
	err = service.AddToCart("user1", "p2")
	if err == nil {
		t.Error("expected error for out of stock")
	}
}

func TestRemoveFromCart(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockCartRepo := mocks.NewMockCartManager(ctrl)
	mockProdRepo := mocks.NewMockProductManager(ctrl)
	mockCouponRepo := mocks.NewMockCouponManager(ctrl)
	service := NewCartService(mockCartRepo, mockProdRepo, mockCouponRepo)

	mockCartRepo.EXPECT().GetCartIDByUserID("user1").Return("cart123", nil)
	mockCartRepo.EXPECT().GetCartItems("cart123").Return([]dto.CartItemsDTO{
		{ProductID: "p1"},
	}, nil)
	mockCartRepo.EXPECT().RemoveFromCart("cart123", "p1").Return(nil)

	err := service.RemoveFromCart("user1", "p1")
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}

	mockCartRepo.EXPECT().GetCartIDByUserID("user2").Return("cart456", nil)
	mockCartRepo.EXPECT().GetCartItems("cart456").Return([]dto.CartItemsDTO{}, nil)
	err = service.RemoveFromCart("user2", "p2")
	if err == nil {
		t.Error("expected error for product not in cart")
	}
}

func TestCheckout(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockCartRepo := mocks.NewMockCartManager(ctrl)
	mockProdRepo := mocks.NewMockProductManager(ctrl)
	mockCouponRepo := mocks.NewMockCouponManager(ctrl)
	service := NewCartService(mockCartRepo, mockProdRepo, mockCouponRepo)

	cartItems := []dto.CartItemsDTO{
		{ProductID: "p1", ProductName: "Item1", Price: 100, Quantity: 2},
	}
	product := models.Product{ID: "p1", Name: "Item1", Stock: 5}

	mockCartRepo.EXPECT().GetCartIDByUserID("user1").Return("cart123", nil)
	mockCartRepo.EXPECT().GetCartItems("cart123").Return(cartItems, nil)
	mockProdRepo.EXPECT().GetProductByID("p1").Return(product, nil)
	mockProdRepo.EXPECT().UpdateProduct(gomock.Any()).Return(nil)
	mockCartRepo.EXPECT().EmptyCart("user1").Return(nil)
	mockCouponRepo.EXPECT().GetCouponByCode("SAVE10").Return(&models.Coupon{Code: "SAVE10", Discount: 10}, nil)

	total, err := service.Checkout("user1", "SAVE10")
	if err != nil || total != 180 {
		t.Errorf("unexpected error or wrong total: %v, total: %v", err, total)
	}

	mockCartRepo.EXPECT().GetCartIDByUserID("user2").Return("cart456", nil)
	mockCartRepo.EXPECT().GetCartItems("cart456").Return(cartItems, nil)
	mockProdRepo.EXPECT().GetProductByID("p1").Return(product, nil)
	mockProdRepo.EXPECT().UpdateProduct(gomock.Any()).Return(nil)
	mockCartRepo.EXPECT().EmptyCart("user2").Return(nil)
	mockCouponRepo.EXPECT().GetCouponByCode("INVALID").Return(nil, errors.New("not found"))

	_, err = service.Checkout("user2", "INVALID")
	if err == nil {
		t.Error("expected error for invalid coupon")
	}
}
