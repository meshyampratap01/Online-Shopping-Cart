package couponRepository

import (
	"database/sql"

	"github.com/meshyampratap01/OnlineShoppingCart/internal/models"
)

type CouponRepository struct {
	db *sql.DB
}

func NewCouponRepository(db *sql.DB) *CouponRepository {
	return &CouponRepository{db: db}
}

func (cr *CouponRepository) SaveCoupon(coupon *models.Coupon) error {
	_, err := cr.db.Exec("INSERT INTO coupons (code, percentage) VALUES (?, ?)",
		coupon.Code, coupon.Percentage)
	return err
}

func (cr *CouponRepository) GetCouponByCode(code string) (*models.Coupon, error) {
	row := cr.db.QueryRow("SELECT code, percentage FROM coupons WHERE code = ?", code)
	coupon := &models.Coupon{}
	err := row.Scan(&coupon.Code, &coupon.Percentage)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}
	return coupon, nil
}

func (cr *CouponRepository) RemoveCoupon(code string) error {
	_, err := cr.db.Exec("DELETE FROM coupons WHERE code = ?", code)
	return err
}
