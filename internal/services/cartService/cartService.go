package cartservice

import (
	"fmt"

	"github.com/meshyampratap01/OnlineShoppingCart/internal/dto"
	cartRepository "github.com/meshyampratap01/OnlineShoppingCart/internal/repository/cartRepository"
	"github.com/meshyampratap01/OnlineShoppingCart/internal/repository/couponRepository"
	"github.com/meshyampratap01/OnlineShoppingCart/internal/repository/productRepository"
)

type CartService struct {
	cartRepo   cartRepository.CartManager
	prodRepo   productRepository.ProductManager
	couponRepo couponRepository.CouponManager
}

func NewCartService(cartRepo cartRepository.CartManager, prodRepo productRepository.ProductManager, couponRepo couponRepository.CouponManager) *CartService {
	return &CartService{cartRepo: cartRepo, prodRepo: prodRepo, couponRepo: couponRepo}
}

func (cs *CartService) GetCartItems(userID string) ([]dto.CartItemsDTO, error) {
	cartID, err := cs.cartRepo.GetCartIDByUserID(userID)
	if err != nil {
		return nil, fmt.Errorf("no cart associated with user,%v", err)
	}
	cartItems, err := cs.cartRepo.GetCartItems(cartID)
	if err != nil {
		return nil, fmt.Errorf("can't fetch cart items: %v", err)
	}
	return cartItems, nil
}

func (cs *CartService) AddToCart(userID, prodID string) error {
	prod, err := cs.prodRepo.GetProductByID(prodID)
	if err != nil {
		return err
	}
	if prod.Stock <= 0 {
		return fmt.Errorf("product %s is out of stock", prod.Name)
	}
	return cs.cartRepo.AddToCart(userID, prod)
}

func (cs *CartService) RemoveFromCart(userID, prodID string) error {
	return cs.cartRepo.RemoveFromCart(userID, prodID)
}

func (cs *CartService) Checkout(userID string, couponCode string) (float32, error) {
	cartID, err := cs.cartRepo.GetCartIDByUserID(userID)
	if err != nil {
		return 0, err
	}
	cartItems, err := cs.cartRepo.GetCartItems(cartID)
	if err != nil {
		return 0, err
	}

	var total float32
	for _, item := range cartItems {
		total += item.Price * float32(item.Quantity)
		prod, err := cs.prodRepo.GetProductByID(item.ProductID)
		if err != nil {
			return 0, fmt.Errorf("product %s not found", item.ProductName)
		}
		if prod.Stock < item.Quantity {
			return 0, fmt.Errorf("insufficient stock for product %s", prod.Name)
		}
		prod.Stock -= item.Quantity
		err = cs.prodRepo.UpdateProduct(prod)
		if err != nil {
			return 0, fmt.Errorf("failed to update stock for product %s", prod.Name)
		}
	}
	err=cs.cartRepo.EmptyCart(userID)
	if err!=nil{
		return 0,fmt.Errorf("can't update cart: %v",err)
	}
	if couponCode != "" {
		fmt.Println("coupon applied")
		coupon, err := cs.couponRepo.GetCouponByCode(couponCode)
		if err != nil {
			return 0, err
		}
		total = total - (total * coupon.Discount / 100)
	}
	return total, nil
}
