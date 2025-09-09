package productService

import "github.com/meshyampratap01/OnlineShoppingCart/internal/models"

//go:generate mockgen -source=interface.go -destination=../../mocks/mock_productService.go -package mocks

type ProductServiceManager interface {
	GetAllProducts() ([]models.Product, error)
	GetProductByID(id string) (models.Product, error)
	GetProductByName(name *string) ([]models.Product, error)
}