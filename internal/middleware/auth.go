package middleware

import (
	"context"
	"net/http"
	"strings"

	"korzadivpn/internal/database"

	"github.com/golang-jwt/jwt/v5"
)

type ContextKey string

var UserEmailKey ContextKey = "userEmail"

func Auth(next http.HandlerFunc) http.HandlerFunc {

	return func(
		w http.ResponseWriter,
		r *http.Request,
	) {

		authHeader := r.Header.Get(
			"Authorization",
		)

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
			func(
				token *jwt.Token,
			) (interface{}, error) {

				if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {

					return nil, jwt.ErrSignatureInvalid
				}

				return SecretKey, nil
			},
		)

		if err != nil || !token.Valid {

			http.Error(
				w,
				"Token invalido",
				http.StatusUnauthorized,
			)

			return
		}

		session, err := database.GetActiveSession(
			tokenString,
		)

		if err != nil || session == nil {

			http.Error(
				w,
				"Sesion no activa",
				http.StatusUnauthorized,
			)

			return
		}

		ctx := context.WithValue(
			r.Context(),
			UserEmailKey,
			session.Email,
		)

		next(
			w,
			r.WithContext(ctx),
		)
	}
}
