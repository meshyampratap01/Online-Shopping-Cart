package adminservice

import (
	"github.com/meshyampratap01/OnlineShoppingCart/internal/repository/couponRepository"
	"github.com/meshyampratap01/OnlineShoppingCart/internal/repository/productRepository"
)

type AdminService struct {
	productRepo productRepository.ProductManager
	couponRepo  couponRepository.CouponManager
}

func NewAdminService(productRepo productRepository.ProductManager, couponRepo couponRepository.CouponManager) AdminServiceManager {
	return &AdminService{
		productRepo: productRepo,
		couponRepo:  couponRepo,
	}
}

func (as *AdminService) AddProduct(name string, price float32, stock int) error {
	// Implementation for adding a product
	return nil
}	

func (as *AdminService) UpdateProduct(id, name string, price float32, stock int) error {
	// Implementation for updating a product
	return nil
}

func (as *AdminService) RemoveProduct(id string) error {
	// Implementation for removing a product
	return nil
}

func (as *AdminService) AddCoupon(code string, percentage float32) error {
	// Implementation for adding a coupon
	return nil
}

func (as *AdminService) RemoveCoupon(code string) error {
	// Implementation for removing a coupon
	return nil
}
