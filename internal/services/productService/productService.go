package productService

import (
	"fmt"

	"github.com/meshyampratap01/OnlineShoppingCart/internal/models"
	"github.com/meshyampratap01/OnlineShoppingCart/internal/repository/productRepository"
	"github.com/meshyampratap01/OnlineShoppingCart/internal/utils"
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

func (ps *ProductService) GetProductByName(name string) ([]models.Product, error) {
	products, err := ps.productRepo.GetProductByName(name)
	if err != nil {
		return nil, fmt.Errorf("no product with specified name found")
	}
	return products, nil
}

func (ps *ProductService) AddProduct(name string, price float32, stock int) (models.Product, error) {
	if name == "" || price <= 0 || stock < 0 {
		return models.Product{}, fmt.Errorf("invalid product details")
	}
	product := models.Product{
		ID:    utils.NewUUID(),
		Name:  name,
		Price: price,
		Stock: stock,
	}
	err := ps.productRepo.AddProduct(product)
	if err != nil {
		return models.Product{}, fmt.Errorf("can't add product %v", err)
	}
	return product, nil
}

func (ps *ProductService) UpdateProduct(id, name string, price float32, stock int) error {
	product, err := ps.productRepo.GetProductByID(id)
	if err != nil {
		return fmt.Errorf("no product with specified id found")
	}
	if name != "" {
		product.Name = name
	}
	if price != 0 {
		product.Price = price
	}
	if stock != 0 {
		product.Stock = stock
	}
	return ps.productRepo.UpdateProduct(product)
}

func (ps *ProductService) RemoveProduct(id string) error {
	err:= ps.productRepo.RemoveProduct(id)
	if err != nil {
		return fmt.Errorf("no product with specified id found")
	}
	return nil
}