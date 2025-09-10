package utils

import (
	"fmt"
	"testing"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"github.com/meshyampratap01/OnlineShoppingCart/internal/config"
	"github.com/meshyampratap01/OnlineShoppingCart/internal/models"
)

func TestNewUUID(t *testing.T) {
	u := NewUUID()

	if u == "" {
		t.Error("wanted a string but got empty")
	}

	_, err := uuid.Parse(u)
	if err != nil {
		t.Errorf("wanted valid uuid, got error: %v", err)
	}
}

func TestHashPassword(t *testing.T) {
	password := "mySecurePassword123"
	hashed, err := HashPassword(password)

	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	if hashed == "" {
		t.Fatal("Expected non-empty hashed password")
	}

	if hashed == password {
		t.Fatal("Hashed password should not be equal to the original password")
	}
}

func TestCheckPassword(t *testing.T) {
	password := "mySecurePassword123"
	wrongPassword := "wrongPassword456"

	hashed, err := HashPassword(password)
	if err != nil {
		t.Fatalf("Hashing failed: %v", err)
	}

	if !CheckPassword(hashed, password) {
		t.Fatal("Expected password to match")
	}

	if CheckPassword(hashed, wrongPassword) {
		t.Fatal("Expected password mismatch")
	}
}

func TestGenerateJWT(t *testing.T) {
	// Set a test secret key
	config.JWT_Secret = []byte("test_secret")

	user := models.UserJWT{
		UserID: "123",
		Email:  "test@example.com",
		Role:   models.Admin,
	}

	tokenStr, err := GenerateJWT(user)
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	if tokenStr == "" {
		t.Fatal("Expected non-empty token string")
	}

	token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method")
		}
		return config.JWT_Secret, nil
	})

	if err != nil || !token.Valid {
		t.Fatalf("Expected valid token, got error: %v", err)
	}

	_, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		t.Fatal("Expected MapClaims")
	}
}
