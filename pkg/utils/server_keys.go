package utils

import (
	"crypto/rand"
	"encoding/base64"
)

func GenerateServerKey() string {

	key := make([]byte, 32)

	_, err := rand.Read(key)

	if err != nil {
		return ""
	}

	return base64.StdEncoding.EncodeToString(key)
}
