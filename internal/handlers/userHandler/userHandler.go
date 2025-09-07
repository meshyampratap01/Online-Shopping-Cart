package userHandler

import (
	"encoding/json"
	"net/http"
	"strings"

	"github.com/meshyampratap01/OnlineShoppingCart/internal/dto"
	"github.com/meshyampratap01/OnlineShoppingCart/internal/models"
	"github.com/meshyampratap01/OnlineShoppingCart/internal/services/userService"
	"github.com/meshyampratap01/OnlineShoppingCart/internal/validators"
	"github.com/meshyampratap01/OnlineShoppingCart/internal/webResponse"
)

type UserHandler struct {
	userService userService.UserServiceManager
}

func NewUserHandler(userService userService.UserServiceManager) *UserHandler {
	return &UserHandler{
		userService: userService,
	}
}

// api/v1/register [POST]
func (uh *UserHandler) RegisterUser(w http.ResponseWriter, r *http.Request) {
	var req dto.SignupRequestDTO
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		resp := webResponse.NewErrorResponse(http.StatusBadRequest, "invalid request body")
		w.WriteHeader(resp.Code)
		json.NewEncoder(w).Encode(resp)
		return
	}

	email := strings.TrimSpace(req.Email)
	email = strings.ToLower(email)
	name := strings.TrimSpace(req.Name)
	password := strings.TrimSpace(req.Password)

	err = validators.ValidateEmail(email)
	if err != nil {
		resp := webResponse.NewErrorResponse(http.StatusBadRequest, err.Error())
		w.WriteHeader(resp.Code)
		json.NewEncoder(w).Encode(resp)
		return
	}

	err = validators.ValidateName(name)
	if err != nil {
		resp := webResponse.NewErrorResponse(http.StatusBadRequest, err.Error())
		w.WriteHeader(resp.Code)
		json.NewEncoder(w).Encode(resp)
		return
	}

	err = validators.ValidatePassword(password)
	if err != nil {
		resp := webResponse.NewErrorResponse(http.StatusBadRequest, err.Error())
		w.WriteHeader(resp.Code)
		json.NewEncoder(w).Encode(resp)
		return
	}

	var role models.UserRole = models.Customer

	err = uh.userService.RegisterUser(name, email, password, role)
	if err != nil {
		resp := webResponse.NewErrorResponse(http.StatusInternalServerError, err.Error())
		w.WriteHeader(resp.Code)
		json.NewEncoder(w).Encode(resp)
		return
	}

	w.WriteHeader(http.StatusCreated)
	reqp := webResponse.NewSuccessResponse(http.StatusCreated, "User registered successfully", nil)
	json.NewEncoder(w).Encode(reqp)
}

// api/v1/login [POST]
func (uh *UserHandler) LoginHandler(w http.ResponseWriter, r *http.Request) {
	var req dto.LoginRequestDTO

	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		resp := webResponse.NewErrorResponse(http.StatusBadRequest, "invalid request body")
		w.WriteHeader(resp.Code)
		json.NewEncoder(w).Encode(resp)
		return
	}

	email := strings.TrimSpace(req.Email)
	email = strings.ToLower(email)

	token, err := uh.userService.Login(email, req.Password)
	if err != nil {
		resp := webResponse.NewErrorResponse(http.StatusUnauthorized, err.Error())
		w.WriteHeader(resp.Code)
		json.NewEncoder(w).Encode(resp)
		return
	}
	resp := webResponse.NewSuccessResponse(http.StatusOK, "Login successful", `{"token":"`+token+`"}`)
	w.WriteHeader(resp.Code)
	json.NewEncoder(w).Encode(resp)
}
