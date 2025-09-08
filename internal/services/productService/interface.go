package productService

import "github.com/meshyampratap01/OnlineShoppingCart/internal/models"

type ProductServiceManager interface {
	GetAllProducts() ([]models.Product, error)
	GetProductByID(id string) (models.Product, error)
	GetProductByName(name *string) ([]models.Product, error)
}