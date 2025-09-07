package utils

import (
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"github.com/meshyampratap01/OnlineShoppingCart/internal/config"
	"github.com/meshyampratap01/OnlineShoppingCart/internal/models"
	"golang.org/x/crypto/bcrypt"
)

func NewUUID() string {
	return uuid.New().String()
}

func GenerateJWT(userJWT models.UserJWT) (string, error) {
	claims := jwt.MapClaims{
		"user_id": userJWT.UserID,
		"email":   userJWT.Email,
		"role":    userJWT.Role,
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(config.JWT_Secret)
	if err != nil {
		return "", err
	}
	return tokenString, nil
}

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(bytes), err
}

func CheckPassword(hashed, password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashed), []byte(password))
	return err == nil
}
