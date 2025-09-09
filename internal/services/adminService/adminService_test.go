package adminservice_test

import (
	"errors"
	"testing"

	"github.com/meshyampratap01/OnlineShoppingCart/internal/mocks"
	"github.com/meshyampratap01/OnlineShoppingCart/internal/models"
	adminservice "github.com/meshyampratap01/OnlineShoppingCart/internal/services/adminService"
	"go.uber.org/mock/gomock"
)

func TestAddProduct(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockProductRepo := mocks.NewMockProductManager(ctrl)
	mockCouponRepo := mocks.NewMockCouponManager(ctrl)
	service := adminservice.NewAdminService(mockProductRepo, mockCouponRepo)

	// Invalid input
	err := service.AddProduct("", 0, -1)
	if err == nil {
		t.Error("expected error for invalid product details")
	}

	// Valid input
	mockProduct := models.Product{Name: "Test", Price: 100, Stock: 10}
	mockProductRepo.EXPECT().AddProduct(gomock.Any()).Return(nil)

	err = service.AddProduct(mockProduct.Name, mockProduct.Price, mockProduct.Stock)
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}
}

func TestUpdateProduct(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockProductRepo := mocks.NewMockProductManager(ctrl)
	mockCouponRepo := mocks.NewMockCouponManager(ctrl)
	service := adminservice.NewAdminService(mockProductRepo, mockCouponRepo)

	product := models.Product{ID: "123", Name: "Old", Price: 50, Stock: 5}
	mockProductRepo.EXPECT().GetProductByID("123").Return(product, nil)
	mockProductRepo.EXPECT().UpdateProduct(gomock.Any()).Return(nil)

	err := service.UpdateProduct("123", "New", 100, 10)
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}

	// Product not found
	mockProductRepo.EXPECT().GetProductByID("404").Return(models.Product{}, errors.New("not found"))
	err = service.UpdateProduct("404", "New", 100, 10)
	if err == nil {
		t.Error("expected error for product not found")
	}
}

func TestRemoveProduct(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockProductRepo := mocks.NewMockProductManager(ctrl)
	mockCouponRepo := mocks.NewMockCouponManager(ctrl)
	service := adminservice.NewAdminService(mockProductRepo, mockCouponRepo)

	product := models.Product{ID: "123"}
	mockProductRepo.EXPECT().GetProductByID("123").Return(product, nil)
	mockProductRepo.EXPECT().RemoveProduct("123").Return(nil)

	err := service.RemoveProduct("123")
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}

	// Product not found
	mockProductRepo.EXPECT().GetProductByID("404").Return(models.Product{}, errors.New("not found"))
	err = service.RemoveProduct("404")
	if err == nil {
		t.Error("expected error for product not found")
	}
}

func TestAddCoupon(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockProductRepo := mocks.NewMockProductManager(ctrl)
	mockCouponRepo := mocks.NewMockCouponManager(ctrl)
	service := adminservice.NewAdminService(mockProductRepo, mockCouponRepo)

	// Invalid coupon
	err := service.AddCoupon("", -10)
	if err == nil {
		t.Error("expected error for invalid coupon")
	}

	// Coupon already exists
	mockCouponRepo.EXPECT().GetCouponByCode("SAVE10").Return(&models.Coupon{Code: "SAVE10"}, nil)
	err = service.AddCoupon("SAVE10", 10)
	if err == nil {
		t.Error("expected error for duplicate coupon")
	}

	// Valid coupon
	mockCouponRepo.EXPECT().GetCouponByCode("NEW10").Return(nil, errors.New("not found"))
	mockCouponRepo.EXPECT().SaveCoupon(gomock.Any()).Return(nil)
	err = service.AddCoupon("NEW10", 10)
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}
}

func TestRemoveCoupon(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockProductRepo := mocks.NewMockProductManager(ctrl)
	mockCouponRepo := mocks.NewMockCouponManager(ctrl)
	service := adminservice.NewAdminService(mockProductRepo, mockCouponRepo)

	// Coupon exists
	mockCouponRepo.EXPECT().GetCouponByCode("SAVE10").Return(&models.Coupon{Code: "SAVE10"}, nil)
	mockCouponRepo.EXPECT().RemoveCoupon("SAVE10").Return(nil)
	err := service.RemoveCoupon("SAVE10")
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}

	// Coupon not found
	mockCouponRepo.EXPECT().GetCouponByCode("INVALID").Return(nil, errors.New("not found"))
	err = service.RemoveCoupon("INVALID")
	if err == nil {
		t.Error("expected error for coupon not found")
	}
}
