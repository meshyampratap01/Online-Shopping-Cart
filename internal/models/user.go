package models

type UserRole int

const (
	Admin UserRole = iota
	Customer
)

type User struct {
	ID       string
	Name     string
	Email    string
	Password string
	Role     UserRole
	Cart     []Product
}
