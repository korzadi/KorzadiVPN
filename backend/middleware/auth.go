package middleware

import (
	"context"
	"net/http"
	"strings"

	"github.com/golang-jwt/jwt/v5"
)

type ContextKey string

var UserEmailKey ContextKey = "userEmail"

func Auth(next http.HandlerFunc) http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {

		authHeader := r.Header.Get("Authorization")

		if authHeader == "" {

			http.Error(
				w,
				"Token requerido",
				http.StatusUnauthorized,
			)

			return
		}

		tokenString := strings.TrimPrefix(
			authHeader,
			"Bearer ",
		)

		token, err := jwt.Parse(
			tokenString,
			func(token *jwt.Token) (interface{}, error) {

				if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {

					return nil, jwt.ErrSignatureInvalid
				}

				return SecretKey, nil
			},
		)

		if err != nil || !token.Valid {

			http.Error(
				w,
				"Token inválido",
				http.StatusUnauthorized,
			)

			return
		}

		claims, ok := token.Claims.(jwt.MapClaims)

		if !ok {

			http.Error(
				w,
				"Claims inválidos",
				http.StatusUnauthorized,
			)

			return
		}

		emailValue, ok := claims["email"].(string)

		if !ok || emailValue == "" {

			http.Error(
				w,
				"Email inválido en token",
				http.StatusUnauthorized,
			)

			return
		}

		ctx := context.WithValue(
			r.Context(),
			UserEmailKey,
			emailValue,
		)

		next(
			w,
			r.WithContext(ctx),
		)

	}

}
