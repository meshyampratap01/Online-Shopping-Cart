package couponRepository

import (
	"database/sql"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/meshyampratap01/OnlineShoppingCart/internal/models"
)

func setupMockDB(t *testing.T) (*sql.DB, sqlmock.Sqlmock, *CouponRepository) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("failed to open sqlmock: %v", err)
	}
	return db, mock, &CouponRepository{db: db}
}


func TestSaveCoupon(t *testing.T) {
	db, mock, repo := setupMockDB(t)
	defer db.Close()

	mock.ExpectExec("INSERT INTO coupons").
		WithArgs("COUPON123", 10.0).
		WillReturnResult(sqlmock.NewResult(1, 1))


	coupon := &models.Coupon{Code: "COUPON123", Discount: 10.0}
	if err := repo.SaveCoupon(coupon); err != nil {
		t.Errorf("unexpected error: %v", err)
	}
}

func TestGetCouponByCode(t *testing.T) {
	db, mock, repo := setupMockDB(t)
	defer db.Close()

	mock.ExpectQuery("SELECT code, discount FROM coupons").
		WithArgs("COUPON123").
		WillReturnRows(sqlmock.NewRows([]string{"code", "discount"}).AddRow("COUPON123", 10.0))

	coupon, err := repo.GetCouponByCode("COUPON123")
	if err != nil || coupon == nil || coupon.Code != "COUPON123" || coupon.Discount != 10.0 {
		t.Errorf("expected COUPON123, 10.0 got %+v, err=%v", coupon, err)
	}

	mock.ExpectQuery("SELECT code, discount FROM coupons").
		WithArgs("INVALID").
		WillReturnError(sql.ErrNoRows)

	coupon, err = repo.GetCouponByCode("INVALID")
	if err == nil || coupon != nil {
		t.Errorf("expected error got %+v, err=%v", coupon, err)
	}	
}

func TestRemoveCoupon(t *testing.T) {
	db, mock, repo := setupMockDB(t)
	defer db.Close()	

	mock.ExpectExec("DELETE FROM coupons").
		WithArgs("COUPON123").
		WillReturnResult(sqlmock.NewResult(1, 1))

	if err := repo.RemoveCoupon("COUPON123"); err != nil {
		t.Errorf("unexpected error: %v", err)
	}
}