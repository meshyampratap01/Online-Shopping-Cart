package adminservice

import (
	"fmt"

	"github.com/meshyampratap01/OnlineShoppingCart/internal/models"
	"github.com/meshyampratap01/OnlineShoppingCart/internal/repository/couponRepository"
	"github.com/meshyampratap01/OnlineShoppingCart/internal/repository/productRepository"
	"github.com/meshyampratap01/OnlineShoppingCart/internal/utils"
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
	if name == "" || price <= 0 || stock < 0 {
		return fmt.Errorf("invalid product details")
	}
	newProduct, err := as.CreateProduct(name, price, stock)
	if err != nil {
		return err
	}
	return as.productRepo.AddProduct(newProduct)
}

func (as *AdminService) CreateProduct(name string, price float32, stock int) (models.Product, error) {
	newProduct := models.Product{
		ID:    utils.NewUUID(),
		Name:  name,
		Price: price,
		Stock: stock,
	}
	return newProduct, nil
}

func (as *AdminService) UpdateProduct(id, name string, price float32, stock int) error {
	product,err := as.productRepo.GetProductByID(id)
	if err != nil {
		return fmt.Errorf("product not found")
	}
	if name != "" {
		product.Name = name
	}
	if price > 0 {
		product.Price = price
	}
	if stock >= 0 {
		product.Stock = stock
	}
	return as.productRepo.UpdateProduct(product)
}

func (as *AdminService) RemoveProduct(id string) error {
	product,err:=as.productRepo.GetProductByID(id)
	if err != nil {
		return fmt.Errorf("product not found")
	}
	return as.productRepo.RemoveProduct(product.ID)
}

func (as *AdminService) AddCoupon(code string, percentage float32) error {
	coupon, err := as.couponRepo.GetCouponByCode(code)
	if err == nil && coupon != nil {
		return fmt.Errorf("coupon code already exists")
	}
	if code == "" || percentage <= 0 || percentage > 100 {
		return fmt.Errorf("invalid coupon details")
	}
	newCoupon := models.Coupon{
		Code:      code,
		Percentage: percentage,
	}
	return as.couponRepo.SaveCoupon(&newCoupon)
}

func (as *AdminService) RemoveCoupon(code string) error {
	coupon, err := as.couponRepo.GetCouponByCode(code)
	if err != nil {
		return fmt.Errorf("coupon not found")
	}
	return as.couponRepo.RemoveCoupon(coupon.Code)
}
