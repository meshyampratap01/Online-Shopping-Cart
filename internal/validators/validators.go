package validators

import (
	"fmt"
	"regexp"

	"github.com/golang-jwt/jwt/v5"
	"github.com/meshyampratap01/OnlineShoppingCart/internal/config"
	"github.com/meshyampratap01/OnlineShoppingCart/internal/models"
)

func ValidateEmail(email string) error {
	regex := regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`)
	if !regex.MatchString(email) {
		return fmt.Errorf("invalid email format, must be in the format 'user@example.com'")
	}
	return nil
}

func ValidatePassword(password string) error {
	if len(password) < 8 {
		return fmt.Errorf("password must be at least 8 characters long")
	}
	letterRegex := regexp.MustCompile(`[a-zA-Z]`)
	numericRegex := regexp.MustCompile(`[\d]`)
	symbolRegex := regexp.MustCompile(`[!@#$%^&*?]`)

	if !letterRegex.MatchString(password) || !numericRegex.MatchString(password) || !symbolRegex.MatchString(password) {
		return fmt.Errorf("password must be alphanumeric with atleast one special character")
	}

	return nil
}

func ValidateName(name string) error {
	if len(name) < 2 {
		return fmt.Errorf("name must be at least 2 characters long")
	}
	if len(name) > 100 {
		return fmt.Errorf("name must be at most 100 characters long")
	}
	return nil
}

func ValidateJWT(tokenStr string) (models.UserJWT, error) {
	if tokenStr == "" {
		return models.UserJWT{}, fmt.Errorf("token is empty")
	}
	var claims models.UserJWT

	token, err := jwt.ParseWithClaims(tokenStr, &claims, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(config.JWT_Secret), nil
	})
	if err != nil {
		return models.UserJWT{}, fmt.Errorf("invalid token: %v", err)
	}
	if !token.Valid {
		return models.UserJWT{}, fmt.Errorf("invalid token")
	}

	return claims, nil
}

func ValidateCoupon(code string, discount float32) error {
	if len(code) < 3 {
		return fmt.Errorf("coupon code must be at least 3 characters long")
	}
	if discount <= 0 || discount > 100 {
		return fmt.Errorf("coupon discount must be between 0 and 100")
	}
	return nil
}
