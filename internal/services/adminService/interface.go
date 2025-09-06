package adminservice


type AdminServiceManager interface {
	AddProduct(name string, price float32, stock int) error
	UpdateProduct(id, name string, price float32, stock int) error
	RemoveProduct(code string) error
	AddCoupon(code string, percentage float32) error
	RemoveCoupon(code string) error
}