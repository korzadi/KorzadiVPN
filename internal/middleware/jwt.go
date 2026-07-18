package middleware

import (
	"time"

	"korzadivpn/config"

	"github.com/golang-jwt/jwt/v5"
)

func GenerateToken(email string) (string, error) {
	now := time.Now()

	token := jwt.NewWithClaims(
		jwt.SigningMethodHS256,
		jwt.MapClaims{
			"email": email,
			"iat":   now.Unix(),
			"exp": now.
				Add(config.TokenDuration).
				Unix(),
		},
	)

	return token.SignedString(config.GetJWTSecret())
}
