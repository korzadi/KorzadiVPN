package middleware

import (
	"time"

	"korzadivpn/config"

	"github.com/golang-jwt/jwt/v5"
)

var SecretKey = []byte(config.JWTSecret)

func GenerateToken(email string) (string, error) {

	token := jwt.NewWithClaims(
		jwt.SigningMethodHS256,
		jwt.MapClaims{
			"email": email,
			"exp": time.Now().
				Add(config.TokenDuration).
				Unix(),
		},
	)

	return token.SignedString(
		SecretKey,
	)
}
