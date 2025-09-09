//go:generate mockgen -source=interface.go -destination=../../mocks/mock_couponRepository.go -package=mocks

package couponRepository

import "github.com/meshyampratap01/OnlineShoppingCart/internal/models"

type CouponManager interface {
	SaveCoupon(*models.Coupon) error
	GetCouponByCode(code string) (*models.Coupon, error)
	RemoveCoupon(code string) error
}