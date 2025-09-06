package app

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	adminhandler "github.com/meshyampratap01/OnlineShoppingCart/internal/handlers/adminHandler"
	"github.com/meshyampratap01/OnlineShoppingCart/internal/handlers/productHandler"
	"github.com/meshyampratap01/OnlineShoppingCart/internal/handlers/userHandler"
	"github.com/meshyampratap01/OnlineShoppingCart/internal/repository/couponRepository"
	"github.com/meshyampratap01/OnlineShoppingCart/internal/repository/productRepository"
	"github.com/meshyampratap01/OnlineShoppingCart/internal/repository/userRepository"
	adminservice "github.com/meshyampratap01/OnlineShoppingCart/internal/services/adminService"
	"github.com/meshyampratap01/OnlineShoppingCart/internal/services/productService"
	"github.com/meshyampratap01/OnlineShoppingCart/internal/services/userService"
)

type App struct {
	db             *sql.DB
	apimux         *http.ServeMux

	UserHandler    userHandler.UserHandler
	ProductHandler productHandler.ProductHandler
	AdminHandler   adminhandler.AdminHandler
}

func NewApp(db *sql.DB) *App {
	userRepo := userRepository.NewUserRepository(db)
	prodRepo := productRepository.NewProductRepository(db)
	couponRepo := couponRepository.NewCouponRepository(db)

	userServ := userService.NewUserService(userRepo, prodRepo, couponRepo)
	prodServ := productService.NewProductService(prodRepo)
	adminServ := adminservice.NewAdminService(prodRepo, couponRepo)

	userHandler := userHandler.NewUserHandler(userServ)
	prodHandler := productHandler.NewProductHandler(prodServ)
	adminHandler := adminhandler.NewAdminHandler(adminServ)

	app := &App{
		db:             db,
		apimux:         http.NewServeMux(),
		UserHandler:    *userHandler,
		ProductHandler: *prodHandler,
		AdminHandler:   *adminHandler,
	}

	app.RegisterRoutes()

	return app
}

func (app *App) Run(){
	fmt.Println("Starting server on :8080")
	err:=http.ListenAndServe(":8080",app.apimux)
	if err!=nil{
		log.Fatalf("Failed to start server: %v",err)
	}
}