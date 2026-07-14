package utils

import (
	"crypto/rand"
	"encoding/base64"

	"golang.org/x/crypto/curve25519"
)

func GenerateWireGuardKeys() (string, string) {

	privateKey := make([]byte, curve25519.ScalarSize)

	_, err := rand.Read(privateKey)

	if err != nil {
		return "", ""
	}

	// Clave privada WireGuard:
	// clamp según X25519
	privateKey[0] &= 248
	privateKey[31] &= 127
	privateKey[31] |= 64

	publicKey, err := curve25519.X25519(
		privateKey,
		curve25519.Basepoint,
	)

	if err != nil {
		return "", ""
	}

	return base64.StdEncoding.EncodeToString(publicKey),
		base64.StdEncoding.EncodeToString(privateKey)
}
