package middleware

import (
	"comments-api/internal/auth"
	"comments-api/utils"
	"context"
	"net/http"
	"strconv"
	"strings"

	"github.com/golang-jwt/jwt/v5"
)

func JWTMiddleware(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        authHeader := r.Header.Get("Authorization")
        if !strings.HasPrefix(authHeader, "Bearer ") {
            utils.SendJSONError(w,"Bearer token needed", http.StatusUnauthorized)
            return
        }

        tokenString := strings.TrimPrefix(authHeader, "Bearer ")

        claims := &auth.CustomClaims{}
        token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
            return []byte("esternocleidomastoideo"), nil
        })

        if err != nil || !token.Valid {
            utils.SendJSONError(w, "Token inválido o expirado", http.StatusUnauthorized)
            return
        }
        idInt, err := strconv.Atoi(claims.UserID)
        if err != nil {
            utils.SendJSONError(w,"Tokens Id invalid", http.StatusUnauthorized)
            return
        }

        ctx := context.WithValue(r.Context(), "user_id", idInt)
        ctx = context.WithValue(ctx, "username", claims.Username)
        next.ServeHTTP(w, r.WithContext(ctx))
    })
}