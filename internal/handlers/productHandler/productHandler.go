package productHandler

import (
	"encoding/json"
	"net/http"

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

// api/v1/products [GET]
func (ph *ProductHandler) GetAllProducts(w http.ResponseWriter, r *http.Request) {
	products, err := ph.productService.GetAllProducts()
	if err != nil {
		resp := webResponse.NewErrorResponse(http.StatusInternalServerError, err.Error())
		w.WriteHeader(resp.Code)
		json.NewEncoder(w).Encode(resp)
		return
	}
	resp := webResponse.NewSuccessResponse(http.StatusOK, "Products retrieved successfully", products)
	w.WriteHeader(resp.Code)
	json.NewEncoder(w).Encode(resp)
}


// api/v1/products/{prodID} [GET]
func (ph *ProductHandler) GetProductByID(w http.ResponseWriter, r *http.Request) {
	prodID:=r.PathValue("prodID")

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


// api/v1/products?name={name} [GET]
func (ph *ProductHandler) GetProductByName(w http.ResponseWriter, r *http.Request){
	name:=r.URL.Query().Get("name")

	products, err := ph.productService.GetProductByName(name)
	if err != nil {
		resp := webResponse.NewErrorResponse(http.StatusInternalServerError, err.Error())
		w.WriteHeader(resp.Code)
		json.NewEncoder(w).Encode(resp)
		return
	}
	resp := webResponse.NewSuccessResponse(http.StatusOK, "Products retrieved successfully", products)
	w.WriteHeader(resp.Code)
	json.NewEncoder(w).Encode(resp)
}