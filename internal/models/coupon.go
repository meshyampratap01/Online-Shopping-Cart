package models


type Coupon struct {
	Code       string  `json:"code"`
	Percentage float32 `json:"percentage"`
}