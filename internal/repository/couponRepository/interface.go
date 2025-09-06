package couponRepository

import "github.com/meshyampratap01/OnlineShoppingCart/internal/models"

type CouponManager interface {
	SaveCoupon(*models.Coupon) error
	GetCouponByCode(code string) (*models.Coupon, error)
}