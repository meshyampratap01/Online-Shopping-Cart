package adminhandler

import (
	"bytes"
	"context"
	"encoding/json"
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

func getAdminContext() context.Context {
	return context.WithValue(context.Background(), config.User, models.UserJWT{Role: models.Admin})
}

func getUserContext() context.Context {
	return context.WithValue(context.Background(), config.User, models.UserJWT{Role: models.Customer})
}

func TestAddProductHandler_Success(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockService := mocks.NewMockAdminServiceManager(ctrl)
	handler := NewAdminHandler(mockService)

	reqBody := dto.ProductDTO{Name: "Laptop", Price: 1000, Stock: 10}
	body, _ := json.Marshal(reqBody)

	req := httptest.NewRequest(http.MethodPost, "/api/v1/admin/product", bytes.NewReader(body))
	req = req.WithContext(getAdminContext())
	w := httptest.NewRecorder()

	mockService.EXPECT().AddProduct("Laptop", float32(1000), 10).Return(nil)

	handler.AddProductHandler(w, req)

	if w.Code != http.StatusCreated {
		t.Errorf("expected status 201, got %d", w.Code)
	}
}

func TestAddProductHandler_InvalidRole(t *testing.T) {
	handler := NewAdminHandler(nil)

	req := httptest.NewRequest(http.MethodPost, "/api/v1/admin/product", nil)
	req = req.WithContext(getUserContext())
	w := httptest.NewRecorder()

	handler.AddProductHandler(w, req)

	if w.Code != http.StatusUnauthorized {
		t.Errorf("expected status 401, got %d", w.Code)
	}
}

func TestUpdateProductHandler_Success(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockService := mocks.NewMockAdminServiceManager(ctrl)
	handler := NewAdminHandler(mockService)

	reqBody := dto.ProductDTO{Name: "Phone", Price: 500, Stock: 5}
	body, _ := json.Marshal(reqBody)

	req := httptest.NewRequest(http.MethodPut, "/api/v1/admin/product/123", bytes.NewReader(body))
	req = req.WithContext(getAdminContext())
	req.SetPathValue("prodID", "123")
	w := httptest.NewRecorder()

	mockService.EXPECT().UpdateProduct("123", "Phone", float32(500), 5).Return(nil)

	handler.UpdateProductHandler(w, req)

	if w.Code != http.StatusOK {
		t.Errorf("expected status 200, got %d", w.Code)
	}
}

func TestRemoveProductHandler_Error(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockService := mocks.NewMockAdminServiceManager(ctrl)
	handler := NewAdminHandler(mockService)

	req := httptest.NewRequest(http.MethodDelete, "/api/v1/admin/product/999", nil)
	req = req.WithContext(getAdminContext())
	req.SetPathValue("prodID", "999")
	w := httptest.NewRecorder()

	mockService.EXPECT().RemoveProduct("999").Return(errors.New("not found"))

	handler.RemoveProductHandler(w, req)

	if w.Code != http.StatusInternalServerError {
		t.Errorf("expected status 500, got %d", w.Code)
	}
}

func TestAddCouponHandler_Success(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockService := mocks.NewMockAdminServiceManager(ctrl)
	handler := NewAdminHandler(mockService)

	reqBody := dto.CouponDTO{Code: "SAVE10", Discount: 10}
	body, _ := json.Marshal(reqBody)

	req := httptest.NewRequest(http.MethodPost, "/api/v1/admin/coupon", bytes.NewReader(body))
	req = req.WithContext(getAdminContext())
	w := httptest.NewRecorder()

	mockService.EXPECT().AddCoupon("SAVE10", float32(10.0)).Return(nil)

	handler.AddCouponHandler(w, req)

	if w.Code != http.StatusCreated {
		t.Errorf("expected status 201, got %d", w.Code)
	}
}

func TestRemoveCouponHandler_Success(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockService := mocks.NewMockAdminServiceManager(ctrl)
	handler := NewAdminHandler(mockService)

	req := httptest.NewRequest(http.MethodDelete, "/api/v1/admin/coupon/SAVE10", nil)
	req = req.WithContext(getAdminContext())
	req.SetPathValue("code", "SAVE10")
	w := httptest.NewRecorder()

	mockService.EXPECT().RemoveCoupon("SAVE10").Return(nil)

	handler.RemoveCouponHandler(w, req)

	if w.Code != http.StatusOK {
		t.Errorf("expected status 200, got %d", w.Code)
	}
}

// Add this to your existing test file

func TestAddProductHandler_InvalidJSON(t *testing.T) {
	handler := NewAdminHandler(nil)

	req := httptest.NewRequest(http.MethodPost, "/api/v1/admin/product", bytes.NewBufferString("{invalid json"))
	req = req.WithContext(getAdminContext())
	w := httptest.NewRecorder()

	handler.AddProductHandler(w, req)

	if w.Code != http.StatusBadRequest {
		t.Errorf("expected 400, got %d", w.Code)
	}
}

func TestAddProductHandler_InvalidName(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockService := mocks.NewMockAdminServiceManager(ctrl)
	handler := NewAdminHandler(mockService)

	reqBody := dto.ProductDTO{Name: "", Price: 100, Stock: 5}
	body, _ := json.Marshal(reqBody)

	req := httptest.NewRequest(http.MethodPost, "/api/v1/admin/product", bytes.NewReader(body))
	req = req.WithContext(getAdminContext())
	w := httptest.NewRecorder()

	handler.AddProductHandler(w, req)

	if w.Code != http.StatusBadRequest {
		t.Errorf("expected 400, got %d", w.Code)
	}
}

func TestAddProductHandler_NegativePrice(t *testing.T) {
	handler := NewAdminHandler(nil)

	reqBody := dto.ProductDTO{Name: "Item", Price: -10, Stock: 5}
	body, _ := json.Marshal(reqBody)

	req := httptest.NewRequest(http.MethodPost, "/api/v1/admin/product", bytes.NewReader(body))
	req = req.WithContext(getAdminContext())
	w := httptest.NewRecorder()

	handler.AddProductHandler(w, req)

	if w.Code != http.StatusBadRequest {
		t.Errorf("expected 400, got %d", w.Code)
	}
}

func TestAddProductHandler_NegativeStock(t *testing.T) {
	handler := NewAdminHandler(nil)

	reqBody := dto.ProductDTO{Name: "Item", Price: 100, Stock: -5}
	body, _ := json.Marshal(reqBody)

	req := httptest.NewRequest(http.MethodPost, "/api/v1/admin/product", bytes.NewReader(body))
	req = req.WithContext(getAdminContext())
	w := httptest.NewRecorder()

	handler.AddProductHandler(w, req)

	if w.Code != http.StatusBadRequest {
		t.Errorf("expected 400, got %d", w.Code)
	}
}

func TestAddProductHandler_ServiceError(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockService := mocks.NewMockAdminServiceManager(ctrl)
	handler := NewAdminHandler(mockService)

	reqBody := dto.ProductDTO{Name: "Item", Price: 100, Stock: 5}
	body, _ := json.Marshal(reqBody)

	req := httptest.NewRequest(http.MethodPost, "/api/v1/admin/product", bytes.NewReader(body))
	req = req.WithContext(getAdminContext())
	w := httptest.NewRecorder()

	mockService.EXPECT().AddProduct("Item", float32(100), 5).Return(errors.New("db error"))

	handler.AddProductHandler(w, req)

	if w.Code != http.StatusInternalServerError {
		t.Errorf("expected 500, got %d", w.Code)
	}
}

func TestUpdateProductHandler_InvalidJSON(t *testing.T) {
	handler := NewAdminHandler(nil)

	req := httptest.NewRequest(http.MethodPut, "/api/v1/admin/product/123", bytes.NewBufferString("{invalid"))
	req = req.WithContext(getAdminContext())
	req.SetPathValue("prodID", "123")
	w := httptest.NewRecorder()

	handler.UpdateProductHandler(w, req)

	if w.Code != http.StatusBadRequest {
		t.Errorf("expected 400, got %d", w.Code)
	}
}

func TestUpdateProductHandler_ServiceError(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockService := mocks.NewMockAdminServiceManager(ctrl)
	handler := NewAdminHandler(mockService)

	reqBody := dto.ProductDTO{Name: "Item", Price: 100, Stock: 5}
	body, _ := json.Marshal(reqBody)

	req := httptest.NewRequest(http.MethodPut, "/api/v1/admin/product/123", bytes.NewReader(body))
	req = req.WithContext(getAdminContext())
	req.SetPathValue("prodID", "123")
	w := httptest.NewRecorder()

	mockService.EXPECT().UpdateProduct("123", "Item", float32(100), 5).Return(errors.New("update failed"))

	handler.UpdateProductHandler(w, req)

	if w.Code != http.StatusInternalServerError {
		t.Errorf("expected 500, got %d", w.Code)
	}
}

func TestAddCouponHandler_InvalidJSON(t *testing.T) {
	handler := NewAdminHandler(nil)

	req := httptest.NewRequest(http.MethodPost, "/api/v1/admin/coupon", bytes.NewBufferString("{invalid"))
	req = req.WithContext(getAdminContext())
	w := httptest.NewRecorder()

	handler.AddCouponHandler(w, req)

	if w.Code != http.StatusBadRequest {
		t.Errorf("expected 400, got %d", w.Code)
	}
}

func TestAddCouponHandler_ValidationError(t *testing.T) {
	handler := NewAdminHandler(nil)

	reqBody := dto.CouponDTO{Code: "", Discount: -5}
	body, _ := json.Marshal(reqBody)

	req := httptest.NewRequest(http.MethodPost, "/api/v1/admin/coupon", bytes.NewReader(body))
	req = req.WithContext(getAdminContext())
	w := httptest.NewRecorder()

	handler.AddCouponHandler(w, req)

	if w.Code != http.StatusBadRequest {
		t.Errorf("expected 400, got %d", w.Code)
	}
}

func TestAddCouponHandler_ServiceError(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockService := mocks.NewMockAdminServiceManager(ctrl)
	handler := NewAdminHandler(mockService)

	reqBody := dto.CouponDTO{Code: "SAVE10", Discount: 10}
	body, _ := json.Marshal(reqBody)

	req := httptest.NewRequest(http.MethodPost, "/api/v1/admin/coupon", bytes.NewReader(body))
	req = req.WithContext(getAdminContext())
	w := httptest.NewRecorder()

	mockService.EXPECT().AddCoupon("SAVE10", float32(10)).Return(errors.New("insert failed"))

	handler.AddCouponHandler(w, req)

	if w.Code != http.StatusInternalServerError {
		t.Errorf("expected 500, got %d", w.Code)
	}
}
