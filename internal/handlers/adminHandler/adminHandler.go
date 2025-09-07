package adminhandler

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/meshyampratap01/OnlineShoppingCart/internal/config"
	"github.com/meshyampratap01/OnlineShoppingCart/internal/dto"
	"github.com/meshyampratap01/OnlineShoppingCart/internal/models"
	adminservice "github.com/meshyampratap01/OnlineShoppingCart/internal/services/adminService"
	"github.com/meshyampratap01/OnlineShoppingCart/internal/validators"
	"github.com/meshyampratap01/OnlineShoppingCart/internal/webResponse"
)

type AdminHandler struct {
	AdminService adminservice.AdminServiceManager
}

func NewAdminHandler(adminService adminservice.AdminServiceManager) *AdminHandler {
	return &AdminHandler{
		AdminService: adminService,
	}
}

// api/v1/admin/product [POST]
func (ah *AdminHandler) AddProductHandler(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	userClaims, ok := ctx.Value(config.User).(models.UserJWT)
	if !ok || userClaims.Role != models.Admin {
		resp := webResponse.NewErrorResponse(http.StatusUnauthorized, "unauthorized")
		w.WriteHeader(resp.Code)
		json.NewEncoder(w).Encode(resp)
		return
	}
	var req dto.ProductDTO
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		resp := webResponse.NewErrorResponse(http.StatusBadRequest, err.Error())
		w.WriteHeader(resp.Code)
		json.NewEncoder(w).Encode(resp)
		return
	}
	err = validators.ValidateName(req.Name)
	if err != nil {
		resp := webResponse.NewErrorResponse(http.StatusBadRequest, fmt.Sprintf("name must be valid: %v",err))
		w.WriteHeader(resp.Code)
		json.NewEncoder(w).Encode(resp)
		return
	}
	if req.Price <= 0 {
		resp := webResponse.NewErrorResponse(http.StatusBadRequest, "price can't be negative or zero")
		w.WriteHeader(resp.Code)
		json.NewEncoder(w).Encode(resp)
		return
	}
	if req.Stock < 0 {
		resp := webResponse.NewErrorResponse(http.StatusBadRequest, "stock can't be negative")
		w.WriteHeader(resp.Code)
		json.NewEncoder(w).Encode(resp)
		return
	}
	err = ah.AdminService.AddProduct(req.Name, req.Price, req.Stock)
	if err != nil {
		resp := webResponse.NewErrorResponse(http.StatusInternalServerError, err.Error())
		w.WriteHeader(resp.Code)
		json.NewEncoder(w).Encode(resp)
		return
	}
	resp := webResponse.NewSuccessResponse(http.StatusCreated, "product added successfully", nil)
	w.WriteHeader(resp.Code)
	json.NewEncoder(w).Encode(resp)
}

// api/v1/admin/product/{prodID} [PUT]
func (ah *AdminHandler) UpdateProductHandler(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	userClaims, ok := ctx.Value(config.User).(models.UserJWT)	
	if !ok || userClaims.Role != models.Admin {
		resp := webResponse.NewErrorResponse(http.StatusUnauthorized, "unauthorized")
		w.WriteHeader(resp.Code)
		json.NewEncoder(w).Encode(resp)
		return
	}
	var req dto.ProductDTO
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		resp := webResponse.NewErrorResponse(http.StatusBadRequest, "invalid request body")
		w.WriteHeader(resp.Code)
		json.NewEncoder(w).Encode(resp)
		return
	}
	err = validators.ValidateName(req.Name)
	if err != nil {
		resp := webResponse.NewErrorResponse(http.StatusBadRequest, err.Error())
		w.WriteHeader(resp.Code)
		json.NewEncoder(w).Encode(resp)
		return
	}
	if req.Price <= 0 {
		resp := webResponse.NewErrorResponse(http.StatusBadRequest, "price can't be negative")
		w.WriteHeader(resp.Code)
		json.NewEncoder(w).Encode(resp)
		return
	}
	if req.Stock <= 0 {
		resp := webResponse.NewErrorResponse(http.StatusBadRequest, "stock can't be negative")
		w.WriteHeader(resp.Code)
		json.NewEncoder(w).Encode(resp)
		return
	}
	prodID := r.PathValue("prodID")
	err = ah.AdminService.UpdateProduct(prodID, req.Name, req.Price, req.Stock)
	if err != nil {
		resp := webResponse.NewErrorResponse(http.StatusInternalServerError, err.Error())
		w.WriteHeader(resp.Code)
		json.NewEncoder(w).Encode(resp)
		return
	}
	resp := webResponse.NewSuccessResponse(http.StatusOK, "product updated successfully", nil)
	w.WriteHeader(resp.Code)
	json.NewEncoder(w).Encode(resp)
}

// api/v1/admin/product/{prodID} [DELETE]
func (ah *AdminHandler) RemoveProductHandler(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	userClaims, ok := ctx.Value(config.User).(models.UserJWT)
	if !ok || userClaims.Role != models.Admin {
		resp := webResponse.NewErrorResponse(http.StatusUnauthorized, "unauthorized")
		w.WriteHeader(resp.Code)
		json.NewEncoder(w).Encode(resp)
		return
	}
	prodID := r.PathValue("prodID")
	err := ah.AdminService.RemoveProduct(prodID)
	if err != nil {
		resp := webResponse.NewErrorResponse(http.StatusInternalServerError, err.Error())
		w.WriteHeader(resp.Code)
		json.NewEncoder(w).Encode(resp)
		return
	}
	resp := webResponse.NewSuccessResponse(http.StatusOK, "product removed successfully", nil)
	w.WriteHeader(resp.Code)
	json.NewEncoder(w).Encode(resp)
}


// api/v1/admin/coupon [POST]
func (ah *AdminHandler) AddCouponHandler(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	userClaims, ok := ctx.Value(config.User).(models.UserJWT)
	if !ok || userClaims.Role != models.Admin {
		resp := webResponse.NewErrorResponse(http.StatusUnauthorized, "unauthorized")
		w.WriteHeader(resp.Code)
		json.NewEncoder(w).Encode(resp)
		return
	}
	var req dto.CouponDTO
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		resp := webResponse.NewErrorResponse(http.StatusBadRequest, "invalid request body")
		w.WriteHeader(resp.Code)
		json.NewEncoder(w).Encode(resp)
		return
	}
	err = validators.ValidateCoupon(req.Code, req.Discount)
	if err != nil {
		resp := webResponse.NewErrorResponse(http.StatusBadRequest, err.Error())
		w.WriteHeader(resp.Code)
		json.NewEncoder(w).Encode(resp)
		return
	}
	err = ah.AdminService.AddCoupon(req.Code, req.Discount)
	if err != nil {
		resp := webResponse.NewErrorResponse(http.StatusInternalServerError, err.Error())
		w.WriteHeader(resp.Code)
		json.NewEncoder(w).Encode(resp)
		return
	}
	resp := webResponse.NewSuccessResponse(http.StatusCreated, "coupon added successfully", nil)
	w.WriteHeader(resp.Code)
	json.NewEncoder(w).Encode(resp)
}

// api/v1/admin/coupon/{code} [DELETE]
func (ah *AdminHandler) RemoveCouponHandler(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	userClaims, ok := ctx.Value(config.User).(models.UserJWT)
	if !ok || userClaims.Role != models.Admin {
		resp := webResponse.NewErrorResponse(http.StatusUnauthorized, "unauthorized")
		w.WriteHeader(resp.Code)
		json.NewEncoder(w).Encode(resp)
		return
	}
	couponCode := r.PathValue("code")
	err := ah.AdminService.RemoveCoupon(couponCode)
	if err != nil {
		resp := webResponse.NewErrorResponse(http.StatusInternalServerError, err.Error())
		w.WriteHeader(resp.Code)
		json.NewEncoder(w).Encode(resp)
		return
	}
	resp := webResponse.NewSuccessResponse(http.StatusOK, "coupon removed successfully", nil)
	w.WriteHeader(resp.Code)
	json.NewEncoder(w).Encode(resp)
}
