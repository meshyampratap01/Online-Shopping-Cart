package productHandler

import (
	"encoding/json"
	"net/http"
	"strings"

	"github.com/meshyampratap01/OnlineShoppingCart/internal/models"
	"github.com/meshyampratap01/OnlineShoppingCart/internal/services/productService"
	"github.com/meshyampratap01/OnlineShoppingCart/internal/webResponse"
)

type ProductHandler struct {
	productService productService.ProductServiceManager
}

func NewProductHandler(productService productService.ProductServiceManager) *ProductHandler {
	return &ProductHandler{
		productService: productService,
	}
}

// api/v1/products [GET] also support "name" query param for searching by name
func (ph *ProductHandler) GetAllProducts(w http.ResponseWriter, r *http.Request) {
	name := r.URL.Query().Get("name")
	name = strings.TrimSpace(name)
	var products []models.Product
	var err error
	if name != "" {
		products, err = ph.productService.GetProductByName(&name)
		if err != nil {
			resp := webResponse.NewErrorResponse(http.StatusInternalServerError, err.Error())
			w.WriteHeader(resp.Code)
			json.NewEncoder(w).Encode(resp)
			return
		}
	}else{
		products, err = ph.productService.GetAllProducts()
		if err != nil {
			resp := webResponse.NewErrorResponse(http.StatusInternalServerError, err.Error())
			w.WriteHeader(resp.Code)
			json.NewEncoder(w).Encode(resp)
			return
		}
	}
	resp := webResponse.NewSuccessResponse(http.StatusOK, "Products retrieved successfully", products)
	w.WriteHeader(resp.Code)
	json.NewEncoder(w).Encode(resp)
}

// api/v1/products/{prodID} [GET]
func (ph *ProductHandler) GetProductByID(w http.ResponseWriter, r *http.Request) {
	prodID := r.PathValue("prodID")

	product, err := ph.productService.GetProductByID(prodID)
	if err != nil {
		resp := webResponse.NewErrorResponse(http.StatusInternalServerError, err.Error())
		w.WriteHeader(resp.Code)
		json.NewEncoder(w).Encode(resp)
		return
	}
	resp := webResponse.NewSuccessResponse(http.StatusOK, "Product retrieved successfully", product)
	w.WriteHeader(resp.Code)
	json.NewEncoder(w).Encode(resp)
}
