package productRepository

import "github.com/meshyampratap01/OnlineShoppingCart/internal/models"

type ProductManager interface {
	AddProduct(models.Product) error
	RemoveProduct(id string) error
	UpdateProduct(models.Product) error
	GetAllProducts() ([]models.Product,error)
	GetProductByName(name string)	([]models.Product,error)
	GetProductByID(id string)	(models.Product,error)
}