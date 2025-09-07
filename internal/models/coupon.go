package models


type Coupon struct {
	Code       string  `json:"code"`
	Discount float32 `json:"discount"`
}