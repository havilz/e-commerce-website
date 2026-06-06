package middleware

import (
	"context"
	"net/http"
	"strings"

	"github.com/golang-jwt/jwt/v5"
	"github.com/havilz/ecommerce-backend/pkg/response"
)

type contextKey string

const (
	ContextUserID contextKey = "userID"
	ContextRole   contextKey = "role"
)

var jwtSecret []byte

func SetJWTSecret(secret string) {
	jwtSecret = []byte(secret)
}

func JWTMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authHeader := r.Header.Get("Authorization")
		if authHeader == "" || !strings.HasPrefix(authHeader, "Bearer ") {
			response.Error(w, http.StatusUnauthorized, "Authorization header missing or invalid")
			return
		}

		tokenStr := strings.TrimPrefix(authHeader, "Bearer ")

		token, err := jwt.Parse(tokenStr, func(t *jwt.Token) (interface{}, error) {
			if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, jwt.ErrSignatureInvalid
			}
			return jwtSecret, nil
		})
		if err != nil || !token.Valid {
			response.Error(w, http.StatusUnauthorized, "Invalid or expired token")
			return
		}

		claims, ok := token.Claims.(jwt.MapClaims)
		if !ok {
			response.Error(w, http.StatusUnauthorized, "Invalid token claims")
			return
		}

		userIDFloat, ok := claims["user_id"].(float64)
		if !ok {
			response.Error(w, http.StatusUnauthorized, "Invalid token payload")
			return
		}

		role, _ := claims["role"].(string)

		ctx := context.WithValue(r.Context(), ContextUserID, uint(userIDFloat))
		ctx = context.WithValue(ctx, ContextRole, role)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func AdminOnly(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		role, ok := r.Context().Value(ContextRole).(string)
		if !ok || role != "admin" {
			response.Error(w, http.StatusForbidden, "Admin access only")
			return
		}
		next.ServeHTTP(w, r)
	})
}

func GetUserID(r *http.Request) uint {
	userID, _ := r.Context().Value(ContextUserID).(uint)
	return userID
}

func GetRole(r *http.Request) string {
	role, _ := r.Context().Value(ContextRole).(string)
	return role
}
