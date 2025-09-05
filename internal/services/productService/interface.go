package productService

import "github.com/meshyampratap01/OnlineShoppingCart/internal/models"

type ProductServiceManager interface {
	GetProductByID(id string) (models.Product,error)
	GetProductByName(name string) ([]models.Product,error)
	AddProduct(name string, price float32, stock int) (*models.Product,error)
	UpdateProduct(id, name string, price float32, stock int)
	RemoveProduct(id string) error
}