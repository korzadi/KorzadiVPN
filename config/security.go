package config

import (
	"fmt"
	"os"
	"strings"
	"time"
)

var JWTSecret string

var TokenDuration = 24 * time.Hour

func init() {
	secret := strings.TrimSpace(os.Getenv("JWT_SECRET"))

	if secret == "" {
		panic("ERROR: JWT_SECRET no configurado. El backend no puede iniciar sin un secreto JWT de producción.")
	}

	if len(secret) < 32 {
		panic("ERROR: JWT_SECRET debe tener mínimo 32 caracteres.")
	}

	JWTSecret = secret
}

func GetJWTSecret() []byte {
	return []byte(JWTSecret)
}

func ValidateSecurityConfig() error {
	if JWTSecret == "" {
		return fmt.Errorf("JWT_SECRET vacío")
	}

	if len(JWTSecret) < 32 {
		return fmt.Errorf("JWT_SECRET demasiado corto")
	}

	return nil
}
