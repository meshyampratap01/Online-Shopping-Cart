package cartHandler

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/meshyampratap01/OnlineShoppingCart/internal/config"
	"github.com/meshyampratap01/OnlineShoppingCart/internal/models"
	cartService "github.com/meshyampratap01/OnlineShoppingCart/internal/services/cartService"
	"github.com/meshyampratap01/OnlineShoppingCart/internal/webResponse"
)

type CartHandler struct {
	cartService cartService.CartServiceManager
}

func NewCartHandler(cartService cartService.CartServiceManager) *CartHandler {
	return &CartHandler{cartService: cartService}
}

// api/v1/cart [GET]
func (ch *CartHandler) GetCartHandler(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	userClaims, ok := ctx.Value(config.User).(models.UserJWT)
	if !ok {
		resp := webResponse.NewErrorResponse(http.StatusUnauthorized, "unauthorized")
		w.WriteHeader(resp.Code)
		json.NewEncoder(w).Encode(resp)
		return
	}
	userRole := userClaims.Role
	if userRole != models.Customer {
		resp := webResponse.NewErrorResponse(http.StatusForbidden, "forbidden")
		w.WriteHeader(resp.Code)
		json.NewEncoder(w).Encode(resp)
		return
	}
	userId := userClaims.UserID
	cartItems, err := ch.cartService.GetCartItems(userId)
	if err != nil {
		resp := webResponse.NewErrorResponse(http.StatusInternalServerError, err.Error())
		w.WriteHeader(resp.Code)
		json.NewEncoder(w).Encode(resp)
		return
	}
	resp := webResponse.NewSuccessResponse(http.StatusOK, "Cart items fetched successfully", cartItems)
	w.WriteHeader(resp.Code)
	json.NewEncoder(w).Encode(resp)
}

// api/v1/cart/{prodID} [POST]
func (ch *CartHandler) AddToCartHandler(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	userClaims, ok := ctx.Value(config.User).(models.UserJWT)
	if !ok {
		resp := webResponse.NewErrorResponse(http.StatusUnauthorized, "unauthorized")
		w.WriteHeader(resp.Code)
		json.NewEncoder(w).Encode(resp)
		return
	}
	userRole := userClaims.Role
	if userRole != models.Customer {
		resp := webResponse.NewErrorResponse(http.StatusForbidden, "forbidden")
		w.WriteHeader(resp.Code)
		json.NewEncoder(w).Encode(resp)
		return
	}
	userId := userClaims.UserID
	prodID := r.PathValue("prodID")
	err := ch.cartService.AddToCart(userId, prodID)
	if err != nil {
		resp := webResponse.NewErrorResponse(http.StatusInternalServerError, err.Error())
		w.WriteHeader(resp.Code)
		json.NewEncoder(w).Encode(resp)
		return
	}
	resp := webResponse.NewSuccessResponse(http.StatusOK, "Product added to cart successfully", nil)
	w.WriteHeader(resp.Code)
	json.NewEncoder(w).Encode(resp)
}

// api/v1/cart/{prodID} [DELETE]
func (ch *CartHandler) RemoveFromCartHandler(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	userClaims, ok := ctx.Value(config.User).(models.UserJWT)
	if !ok {
		resp := webResponse.NewErrorResponse(http.StatusUnauthorized, "unauthorized")
		w.WriteHeader(resp.Code)
		json.NewEncoder(w).Encode(resp)
		return
	}
	userRole := userClaims.Role
	if userRole != models.Customer {
		resp := webResponse.NewErrorResponse(http.StatusForbidden, "forbidden")
		w.WriteHeader(resp.Code)
		json.NewEncoder(w).Encode(resp)
		return
	}
	userId := userClaims.UserID
	prodID := r.PathValue("prodID")
	err := ch.cartService.RemoveFromCart(userId, prodID)
	if err != nil {
		resp := webResponse.NewErrorResponse(http.StatusInternalServerError, err.Error())
		w.WriteHeader(resp.Code)
		json.NewEncoder(w).Encode(resp)
		return
	}
	resp := webResponse.NewSuccessResponse(http.StatusOK, "Product removed from cart successfully", nil)
	w.WriteHeader(resp.Code)
	json.NewEncoder(w).Encode(resp)
}

// api/v1/checkout [POST]
func (ch *CartHandler) CheckOutHandler(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	userClaims, ok := ctx.Value(config.User).(models.UserJWT)
	if !ok {
		resp := webResponse.NewErrorResponse(http.StatusUnauthorized, "unauthorized")
		w.WriteHeader(resp.Code)
		json.NewEncoder(w).Encode(resp)
		return
	}
	userRole := userClaims.Role
	if userRole != models.Customer {
		resp := webResponse.NewErrorResponse(http.StatusForbidden, "forbidden")
		w.WriteHeader(resp.Code)
		json.NewEncoder(w).Encode(resp)
		return
	}
	userId := userClaims.UserID
	couponCode := r.URL.Query().Get("coupon")
	finalAmount, err := ch.cartService.Checkout(userId, couponCode)
	if err != nil {
		resp := webResponse.NewErrorResponse(http.StatusInternalServerError, err.Error())
		w.WriteHeader(resp.Code)
		json.NewEncoder(w).Encode(resp)
		return
	}
	resp := webResponse.NewSuccessResponse(http.StatusOK, "Checkout successful", `Final Amount: `+fmt.Sprintf("%.2f", finalAmount))
	w.WriteHeader(resp.Code)
	json.NewEncoder(w).Encode(resp)
}
