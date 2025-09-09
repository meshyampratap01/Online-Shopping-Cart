package adminservice

//go:generate mockgen -source=interface.go -destination=../../mocks/mock_adminServcie.go -package mocks


type AdminServiceManager interface {
	AddProduct(name string, price float32, stock int) error
	UpdateProduct(id, name string, price float32, stock int) error
	RemoveProduct(code string) error
	AddCoupon(code string, discount float32) error
	RemoveCoupon(code string) error
}