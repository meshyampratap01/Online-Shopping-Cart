package dto

type CouponDTO struct {
	Code       string  `json:"code"`
	Discount float32 `json:"discount"`
}