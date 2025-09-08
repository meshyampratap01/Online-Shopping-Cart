package productService

import (
	"fmt"

	"github.com/meshyampratap01/OnlineShoppingCart/internal/models"
	"github.com/meshyampratap01/OnlineShoppingCart/internal/repository/productRepository"
)

type ProductService struct {
	productRepo productRepository.ProductManager
}

func NewProductService(productRepo productRepository.ProductManager) ProductServiceManager {
	return &ProductService{productRepo: productRepo}
}

func (ps *ProductService) GetAllProducts() ([]models.Product, error) {
	products, err := ps.productRepo.GetAllProducts()
	if err != nil {
		return nil, fmt.Errorf("can not fetch products")
	}
	return products, nil
}

func (ps *ProductService) GetProductByID(id string) (models.Product, error) {
	product, err := ps.productRepo.GetProductByID(id)
	if err != nil {
		return models.Product{}, fmt.Errorf("no product with specified id found")
	}
	return product, nil
}

func (ps *ProductService) GetProductByName(name *string) ([]models.Product, error) {
	products, err := ps.productRepo.GetProductByName(name)
	if err != nil {
		return nil, fmt.Errorf("no product with specified name found")
	}
	return products, nil
}
