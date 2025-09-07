package models

import "github.com/golang-jwt/jwt/v5"

type UserRole int

const (
	Admin UserRole = iota
	Customer
)

func (ur UserRole) String() string {
	switch ur {
	case Admin:
		return "Admin"
	case Customer:
		return "Customer"
	}
	return "Unknown"
}

type User struct {
	ID       string `json:"id"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
	Role     UserRole `json:"role"`
}

type UserJWT struct {
	UserID string   `json:"user_id"`
	Email  string   `json:"email"`
	Role   UserRole `json:"role"`
	jwt.RegisteredClaims
}