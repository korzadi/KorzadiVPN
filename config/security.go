package config

import (
	"log"
	"os"
	"time"
)

var JWTSecret = loadRequiredEnv("JWT_SECRET")

var TokenDuration = 24 * time.Hour

func loadRequiredEnv(key string) string {
	value := os.Getenv(key)
	if value == "" {
		log.Fatalf("Error crítico: Variable de entorno %s no configurada", key)
	}
	if len(value) < 32 {
		log.Fatalf("Error crítico: Variable de entorno %s debe tener al menos 32 caracteres", key)
	}
	return value
}
