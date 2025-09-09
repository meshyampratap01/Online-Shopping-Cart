package productHandler

import (
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/meshyampratap01/OnlineShoppingCart/internal/mocks"
	"github.com/meshyampratap01/OnlineShoppingCart/internal/models"
	"go.uber.org/mock/gomock"
)

func TestGetAllProducts_NoQuery_Success(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockProductService := mocks.NewMockProductServiceManager(ctrl)
	handler := NewProductHandler(mockProductService)

	req := httptest.NewRequest(http.MethodGet, "/api/v1/products", nil)
	w := httptest.NewRecorder()

	mockProductService.EXPECT().GetAllProducts().Return([]models.Product{
		{ID: "p1", Name: "Laptop", Price: 1000, Stock: 10},
	}, nil)

	handler.GetAllProducts(w, req)

	if w.Code != http.StatusOK {
		t.Errorf("expected 200, got %d", w.Code)
	}
}

func TestGetAllProducts_WithQuery_Success(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockProductService := mocks.NewMockProductServiceManager(ctrl)
	handler := NewProductHandler(mockProductService)

	req := httptest.NewRequest(http.MethodGet, "/api/v1/products?name=Laptop", nil)
	w := httptest.NewRecorder()

	name := "Laptop"
	mockProductService.EXPECT().GetProductByName(&name).Return([]models.Product{
		{ID: "p1", Name: "Laptop", Price: 1000, Stock: 10},
	}, nil)

	handler.GetAllProducts(w, req)

	if w.Code != http.StatusOK {
		t.Errorf("expected 200, got %d", w.Code)
	}
}

func TestGetAllProducts_WithQuery_Error(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockProductService := mocks.NewMockProductServiceManager(ctrl)
	handler := NewProductHandler(mockProductService)

	req := httptest.NewRequest(http.MethodGet, "/api/v1/products?name=Phone", nil)
	w := httptest.NewRecorder()

	name := "Phone"
	mockProductService.EXPECT().GetProductByName(&name).Return(nil, errors.New("db error"))

	handler.GetAllProducts(w, req)

	if w.Code != http.StatusInternalServerError {
		t.Errorf("expected 500, got %d", w.Code)
	}
}

func TestGetAllProducts_NoQuery_Error(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockProductService := mocks.NewMockProductServiceManager(ctrl)
	handler := NewProductHandler(mockProductService)

	req := httptest.NewRequest(http.MethodGet, "/api/v1/products", nil)
	w := httptest.NewRecorder()

	mockProductService.EXPECT().GetAllProducts().Return(nil, errors.New("db error"))

	handler.GetAllProducts(w, req)

	if w.Code != http.StatusInternalServerError {
		t.Errorf("expected 500, got %d", w.Code)
	}
}

func TestGetProductByID_Success(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockProductService := mocks.NewMockProductServiceManager(ctrl)
	handler := NewProductHandler(mockProductService)

	req := httptest.NewRequest(http.MethodGet, "/api/v1/products/p1", nil)
	req.SetPathValue("prodID", "p1")
	w := httptest.NewRecorder()

	mockProductService.EXPECT().GetProductByID("p1").Return(models.Product{
		ID: "p1", Name: "Laptop", Price: 1000, Stock: 10,
	}, nil)

	handler.GetProductByID(w, req)

	if w.Code != http.StatusOK {
		t.Errorf("expected 200, got %d", w.Code)
	}
}

func TestGetProductByID_Error(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockProductService := mocks.NewMockProductServiceManager(ctrl)
	handler := NewProductHandler(mockProductService)

	req := httptest.NewRequest(http.MethodGet, "/api/v1/products/p1", nil)
	req.SetPathValue("prodID", "p1")
	w := httptest.NewRecorder()

	mockProductService.EXPECT().GetProductByID("p1").Return(models.Product{}, errors.New("not found"))

	handler.GetProductByID(w, req)

	if w.Code != http.StatusInternalServerError {
		t.Errorf("expected 500, got %d", w.Code)
	}
}
