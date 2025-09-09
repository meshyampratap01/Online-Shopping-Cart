package userHandler

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/meshyampratap01/OnlineShoppingCart/internal/dto"
	"github.com/meshyampratap01/OnlineShoppingCart/internal/mocks"
	"github.com/meshyampratap01/OnlineShoppingCart/internal/models"
	"go.uber.org/mock/gomock"
)

func TestRegisterUser_Success(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockUserService := mocks.NewMockUserServiceManager(ctrl)
	handler := NewUserHandler(mockUserService)

	reqBody := dto.SignupRequestDTO{
		Name:     "Shyam",
		Email:    "shyam@example.com",
		Password: "StrongPass@123",
	}
	body, _ := json.Marshal(reqBody)

	req := httptest.NewRequest(http.MethodPost, "/api/v1/register", bytes.NewReader(body))
	w := httptest.NewRecorder()

	mockUserService.EXPECT().RegisterUser("Shyam", "shyam@example.com", "StrongPass@123", models.Customer).Return(nil)

	handler.RegisterUser(w, req)

	if w.Code != http.StatusCreated {
		t.Errorf("expected 201, got %d", w.Code)
	}
}


func TestRegisterUser_InvalidEmail(t *testing.T) {
	handler := NewUserHandler(nil)

	reqBody := dto.SignupRequestDTO{
		Name:     "Shyam",
		Email:    "invalid-email",
		Password: "StrongPass@123",
	}
	body, _ := json.Marshal(reqBody)

	req := httptest.NewRequest(http.MethodPost, "/api/v1/register", bytes.NewReader(body))
	w := httptest.NewRecorder()

	handler.RegisterUser(w, req)

	if w.Code != http.StatusBadRequest {
		t.Errorf("expected 400, got %d", w.Code)
	}
}

func TestRegisterUser_InvalidName(t *testing.T) {
	handler := NewUserHandler(nil)

	reqBody := dto.SignupRequestDTO{
		Name:     "",
		Email:    "shyam@example.com",
		Password: "StrongPass@123",
	}
	body, _ := json.Marshal(reqBody)

	req := httptest.NewRequest(http.MethodPost, "/api/v1/register", bytes.NewReader(body))
	w := httptest.NewRecorder()

	handler.RegisterUser(w, req)

	if w.Code != http.StatusBadRequest {
		t.Errorf("expected 400, got %d", w.Code)
	}
}

func TestRegisterUser_InvalidPassword(t *testing.T) {
	handler := NewUserHandler(nil)

	reqBody := dto.SignupRequestDTO{
		Name:     "Shyam",
		Email:    "shyam@example.com",
		Password: "123",
	}
	body, _ := json.Marshal(reqBody)

	req := httptest.NewRequest(http.MethodPost, "/api/v1/register", bytes.NewReader(body))
	w := httptest.NewRecorder()

	handler.RegisterUser(w, req)

	if w.Code != http.StatusBadRequest {
		t.Errorf("expected 400, got %d", w.Code)
	}
}

func TestRegisterUser_ServiceError(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockUserService := mocks.NewMockUserServiceManager(ctrl)
	handler := NewUserHandler(mockUserService)

	reqBody := dto.SignupRequestDTO{
		Name:     "Shyam",
		Email:    "shyam@example.com",
		Password: "StrongPass@123",
	}
	body, _ := json.Marshal(reqBody)

	req := httptest.NewRequest(http.MethodPost, "/api/v1/register", bytes.NewReader(body))
	w := httptest.NewRecorder()

	mockUserService.EXPECT().RegisterUser("Shyam", "shyam@example.com", "StrongPass@123", models.Customer).Return(errors.New("db error"))

	handler.RegisterUser(w, req)

	if w.Code != http.StatusInternalServerError {
		t.Errorf("expected 500, got %d", w.Code)
	}
}

func TestLoginHandler_Success(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockUserService := mocks.NewMockUserServiceManager(ctrl)
	handler := NewUserHandler(mockUserService)

	reqBody := dto.LoginRequestDTO{
		Email:    "shyam@example.com",
		Password: "StrongPass@123",
	}
	body, _ := json.Marshal(reqBody)

	req := httptest.NewRequest(http.MethodPost, "/api/v1/login", bytes.NewReader(body))
	w := httptest.NewRecorder()

	mockUserService.EXPECT().Login("shyam@example.com", "StrongPass@123").Return("token123", nil)

	handler.LoginHandler(w, req)

	if w.Code != http.StatusOK {
		t.Errorf("expected 200, got %d", w.Code)
	}
}

func TestLoginHandler_Unauthorized(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockUserService := mocks.NewMockUserServiceManager(ctrl)
	handler := NewUserHandler(mockUserService)

	reqBody := dto.LoginRequestDTO{
		Email:    "shyam@example.com",
		Password: "wrongpass",
	}
	body, _ := json.Marshal(reqBody)

	req := httptest.NewRequest(http.MethodPost, "/api/v1/login", bytes.NewReader(body))
	w := httptest.NewRecorder()

	mockUserService.EXPECT().Login("shyam@example.com", "wrongpass").Return("", errors.New("invalid credentials"))

	handler.LoginHandler(w, req)

	if w.Code != http.StatusUnauthorized {
		t.Errorf("expected 401, got %d", w.Code)
	}
}
