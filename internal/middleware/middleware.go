package middleware

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/meshyampratap01/OnlineShoppingCart/internal/config"
	"github.com/meshyampratap01/OnlineShoppingCart/internal/models"
	"github.com/meshyampratap01/OnlineShoppingCart/internal/validators"
	"github.com/meshyampratap01/OnlineShoppingCart/internal/webResponse"
)

func AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authHeader := r.Header.Get("Authorization")
		if authHeader == "" {
			resp := webResponse.NewErrorResponse(http.StatusUnauthorized, "missing authorization header")
			w.WriteHeader(resp.Code)
			json.NewEncoder(w).Encode(resp)
			return
		}
		tokenStr := authHeader[len("Bearer "):]

		var claims models.UserJWT

		claims, err := validators.ValidateJWT(tokenStr)
		if err != nil {
			resp := webResponse.NewErrorResponse(http.StatusUnauthorized, err.Error())
			w.WriteHeader(resp.Code)
			json.NewEncoder(w).Encode(resp)
			return
		}

		ctx := context.WithValue(r.Context(), config.User, claims)

		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
