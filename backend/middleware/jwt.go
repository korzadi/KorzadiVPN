package middleware

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
)

var SecretKey = []byte("KorzadiVPN_SECRET_KEY")

func GenerateToken(email string) (string, error) {

	token := jwt.NewWithClaims(jwt.SigningMethodHS256,
		jwt.MapClaims{
			"email": email,
			"exp":   time.Now().Add(time.Hour * 24).Unix(),
		})

	return token.SignedString(SecretKey)
}
