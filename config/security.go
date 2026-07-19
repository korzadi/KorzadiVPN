package config

import (
	"os"
	"time"
)

var JWTSecret = getEnv("JWT_SECRET", "KorzadiVPN-Secret-Desarrollo-Cambiar")

var TokenDuration = 24 * time.Hour

func getEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}
