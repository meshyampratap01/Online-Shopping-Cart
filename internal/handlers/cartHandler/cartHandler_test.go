package cartHandler

import (
	"context"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/meshyampratap01/OnlineShoppingCart/internal/config"
	"github.com/meshyampratap01/OnlineShoppingCart/internal/dto"
	"github.com/meshyampratap01/OnlineShoppingCart/internal/mocks"
	"github.com/meshyampratap01/OnlineShoppingCart/internal/models"
	"go.uber.org/mock/gomock"
)

func getCustomerContext() context.Context {
	return context.WithValue(context.Background(), config.User, models.UserJWT{Role: models.Customer, UserID: "user123"})
}

func getAdminContext() context.Context {
	return context.WithValue(context.Background(), config.User, models.UserJWT{Role: models.Admin, UserID: "admin123"})
}

func TestGetCartHandler_Success(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockCartService := mocks.NewMockCartServiceManager(ctrl)
	handler := NewCartHandler(mockCartService)

	req := httptest.NewRequest(http.MethodGet, "/api/v1/cart", nil)
	req = req.WithContext(getCustomerContext())
	w := httptest.NewRecorder()

	mockCartService.EXPECT().GetCartItems("user123").Return(nil, nil)

	handler.GetCartHandler(w, req)

	if w.Code != http.StatusOK {
		t.Errorf("expected 200, got %d", w.Code)
	}
}

func TestGetCartHandler_EmptyCart(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockCartService := mocks.NewMockCartServiceManager(ctrl)
	handler := NewCartHandler(mockCartService)

	req := httptest.NewRequest(http.MethodGet, "/api/v1/cart", nil)
	req = req.WithContext(getCustomerContext())
	w := httptest.NewRecorder()

	mockCartService.EXPECT().GetCartItems("user123").Return([]dto.CartItemsDTO{}, nil)

	handler.GetCartHandler(w, req)

	if w.Code != http.StatusOK {
		t.Errorf("expected 200, got %d", w.Code)
	}
}

func TestGetCartHandler_Unauthorized(t *testing.T) {
	handler := NewCartHandler(nil)

	req := httptest.NewRequest(http.MethodGet, "/api/v1/cart", nil)
	w := httptest.NewRecorder()

	handler.GetCartHandler(w, req)

	if w.Code != http.StatusUnauthorized {
		t.Errorf("expected 401, got %d", w.Code)
	}
}

func TestGetCartHandler_Forbidden(t *testing.T) {
	handler := NewCartHandler(nil)

	req := httptest.NewRequest(http.MethodGet, "/api/v1/cart", nil)
	req = req.WithContext(getAdminContext())
	w := httptest.NewRecorder()

	handler.GetCartHandler(w, req)

	if w.Code != http.StatusForbidden {
		t.Errorf("expected 403, got %d", w.Code)
	}
}

func TestGetCartHandler_ServiceError(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockCartService := mocks.NewMockCartServiceManager(ctrl)
	handler := NewCartHandler(mockCartService)

	req := httptest.NewRequest(http.MethodGet, "/api/v1/cart", nil)
	req = req.WithContext(getCustomerContext())
	w := httptest.NewRecorder()

	mockCartService.EXPECT().GetCartItems("user123").Return(nil, errors.New("db error"))

	handler.GetCartHandler(w, req)

	if w.Code != http.StatusInternalServerError {
		t.Errorf("expected 500, got %d", w.Code)
	}
}

func TestAddToCartHandler_Success(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockCartService := mocks.NewMockCartServiceManager(ctrl)
	handler := NewCartHandler(mockCartService)

	req := httptest.NewRequest(http.MethodPost, "/api/v1/cart/p1", nil)
	req = req.WithContext(getCustomerContext())
	req.SetPathValue("prodID", "p1")
	w := httptest.NewRecorder()

	mockCartService.EXPECT().AddToCart("user123", "p1").Return(nil)

	handler.AddToCartHandler(w, req)

	if w.Code != http.StatusOK {
		t.Errorf("expected 200, got %d", w.Code)
	}
}

func TestAddToCartHandler_Error(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockCartService := mocks.NewMockCartServiceManager(ctrl)
	handler := NewCartHandler(mockCartService)

	req := httptest.NewRequest(http.MethodPost, "/api/v1/cart/p1", nil)
	req = req.WithContext(getCustomerContext())
	req.SetPathValue("prodID", "p1")
	w := httptest.NewRecorder()

	mockCartService.EXPECT().AddToCart("user123", "p1").Return(errors.New("add error"))

	handler.AddToCartHandler(w, req)

	if w.Code != http.StatusInternalServerError {
		t.Errorf("expected 500, got %d", w.Code)
	}
}

func TestRemoveFromCartHandler_Success(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockCartService := mocks.NewMockCartServiceManager(ctrl)
	handler := NewCartHandler(mockCartService)

	req := httptest.NewRequest(http.MethodDelete, "/api/v1/cart/p1", nil)
	req = req.WithContext(getCustomerContext())
	req.SetPathValue("prodID", "p1")
	w := httptest.NewRecorder()

	mockCartService.EXPECT().RemoveFromCart("user123", "p1").Return(nil)

	handler.RemoveFromCartHandler(w, req)

	if w.Code != http.StatusOK {
		t.Errorf("expected 200, got %d", w.Code)
	}
}

func TestRemoveFromCartHandler_Error(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockCartService := mocks.NewMockCartServiceManager(ctrl)
	handler := NewCartHandler(mockCartService)

	req := httptest.NewRequest(http.MethodDelete, "/api/v1/cart/p1", nil)
	req = req.WithContext(getCustomerContext())
	req.SetPathValue("prodID", "p1")
	w := httptest.NewRecorder()

	mockCartService.EXPECT().RemoveFromCart("user123", "p1").Return(errors.New("remove error"))

	handler.RemoveFromCartHandler(w, req)

	if w.Code != http.StatusInternalServerError {
		t.Errorf("expected 500, got %d", w.Code)
	}
}

func TestCheckOutHandler_Success(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockCartService := mocks.NewMockCartServiceManager(ctrl)
	handler := NewCartHandler(mockCartService)

	req := httptest.NewRequest(http.MethodPost, "/api/v1/checkout?code=SAVE10", nil)
	req = req.WithContext(getCustomerContext())
	w := httptest.NewRecorder()

	mockCartService.EXPECT().Checkout("user123", "SAVE10").Return(float32(250.0), nil)

	handler.CheckOutHandler(w, req)

	if w.Code != http.StatusOK {
		t.Errorf("expected 200, got %d", w.Code)
	}
}

func TestCheckOutHandler_Error(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockCartService := mocks.NewMockCartServiceManager(ctrl)
	handler := NewCartHandler(mockCartService)

	req := httptest.NewRequest(http.MethodPost, "/api/v1/checkout?code=SAVE10", nil)
	req = req.WithContext(getCustomerContext())
	w := httptest.NewRecorder()

	mockCartService.EXPECT().Checkout("user123", "SAVE10").Return(float32(0.0), errors.New("checkout failed"))

	handler.CheckOutHandler(w, req)

	if w.Code != http.StatusInternalServerError {
		t.Errorf("expected 500, got %d", w.Code)
	}
}
