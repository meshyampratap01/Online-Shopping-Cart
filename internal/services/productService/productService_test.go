package productService

import (
	"errors"
	"testing"

	"github.com/meshyampratap01/OnlineShoppingCart/internal/mocks"
	"github.com/meshyampratap01/OnlineShoppingCart/internal/models"
	"go.uber.org/mock/gomock"
)

func TestGetAllProducts(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mocks.NewMockProductManager(ctrl)
	service := NewProductService(mockRepo)

	expectedProducts := []models.Product{
		{ID: "1", Name: "Product1", Price: 100, Stock: 10},
		{ID: "2", Name: "Product2", Price: 200, Stock: 5},
	}

	mockRepo.EXPECT().GetAllProducts().Return(expectedProducts, nil)

	products, err := service.GetAllProducts()
	if err != nil || len(products) != 2 {
		t.Errorf("unexpected error or wrong product count: %v", err)
	}

	mockRepo.EXPECT().GetAllProducts().Return(nil, errors.New("db error"))
	_, err = service.GetAllProducts()
	if err == nil {
		t.Error("expected error for failed fetch")
	}
}

func TestGetProductByID(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mocks.NewMockProductManager(ctrl)
	service := NewProductService(mockRepo)

	expectedProduct := models.Product{ID: "1", Name: "Product1", Price: 100, Stock: 10}
	mockRepo.EXPECT().GetProductByID("1").Return(expectedProduct, nil)

	product, err := service.GetProductByID("1")
	if err != nil || product.ID != "1" {
		t.Errorf("unexpected error or wrong product: %v", err)
	}

	mockRepo.EXPECT().GetProductByID("404").Return(models.Product{}, errors.New("not found"))
	_, err = service.GetProductByID("404")
	if err == nil {
		t.Error("expected error for missing product")
	}
}

func TestGetProductByName(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mocks.NewMockProductManager(ctrl)
	service := NewProductService(mockRepo)

	name := "Product1"
	expectedProducts := []models.Product{
		{ID: "1", Name: "Product1", Price: 100, Stock: 10},
	}

	mockRepo.EXPECT().GetProductByName(&name).Return(expectedProducts, nil)

	products, err := service.GetProductByName(&name)
	if err != nil || len(products) != 1 {
		t.Errorf("unexpected error or wrong product count: %v", err)
	}

	mockRepo.EXPECT().GetProductByName(&name).Return(nil, errors.New("not found"))
	_, err = service.GetProductByName(&name)
	if err == nil {
		t.Error("expected error for missing product by name")
	}
}
