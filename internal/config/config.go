package config

type ContextKey string

const (
	User ContextKey = "user"
)

var (
	JWT_Secret = []byte("my_jwt_secret_key") // In production, use a secure method to manage secrets
)