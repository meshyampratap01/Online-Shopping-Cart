package app

import (
	"net/http"

	"github.com/meshyampratap01/OnlineShoppingCart/internal/middleware"
)

var baseURL = "/api/v1/"

func withAuth(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		middleware.AuthMiddleware(next).ServeHTTP(w, r)
	}
}


func (app *App) RegisterRoutes() {
	app.apimux.HandleFunc("POST "+baseURL+"/register", app.UserHandler.RegisterUser)
	app.apimux.HandleFunc("POST "+baseURL+"/login", app.UserHandler.LoginHandler)

	app.apimux.HandleFunc("GET "+baseURL+"/products", app.ProductHandler.GetAllProducts)
	app.apimux.HandleFunc("GET "+baseURL+"/products/{prodID}", app.ProductHandler.GetProductByID)
	app.apimux.HandleFunc("GET "+baseURL+"/products/", app.ProductHandler.GetProductByName) // Search by name query param

	app.apimux.HandleFunc("POST "+baseURL+"/cart/{prodID}", withAuth(app.CartHandler.AddToCartHandler))
	app.apimux.HandleFunc("GET "+baseURL+"/cart", withAuth(app.CartHandler.GetCartHandler))
	app.apimux.HandleFunc("DELETE "+baseURL+"/cart/{prodID}", withAuth(app.CartHandler.RemoveFromCartHandler))
	app.apimux.HandleFunc("POST "+baseURL+"/checkout", withAuth(app.CartHandler.CheckOutHandler))

	app.apimux.HandleFunc("POST "+baseURL+"/admin/products", withAuth(app.AdminHandler.AddProductHandler))
	app.apimux.HandleFunc("PUT "+baseURL+"/admin/products/{prodID}", withAuth(app.AdminHandler.UpdateProductHandler))
	app.apimux.HandleFunc("DELETE "+baseURL+"/admin/products/{prodID}", withAuth(app.AdminHandler.RemoveProductHandler))

	app.apimux.HandleFunc("POST "+baseURL+"/admin/coupons", withAuth(app.AdminHandler.AddCouponHandler))
	app.apimux.HandleFunc("DELETE "+baseURL+"/admin/coupons/{code}", withAuth(app.AdminHandler.RemoveCouponHandler))
}


