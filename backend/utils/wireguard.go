package utils

import (
	"crypto/rand"
	"encoding/base64"
)

func GenerateKey() string {

	key := make([]byte, 32)

	_, err := rand.Read(key)

	if err != nil {
		return ""
	}

	return base64.StdEncoding.EncodeToString(key)
}

func GenerateWireGuardKeys() (string, string) {

	privateKey := GenerateKey()

	publicKey := GenerateKey()

	return publicKey, privateKey
}
