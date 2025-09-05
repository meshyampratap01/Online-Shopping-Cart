package productService

import "github.com/meshyampratap01/OnlineShoppingCart/internal/repository/productRepository"

type ProductService struct {
	productRepo productRepository.ProductManager
}