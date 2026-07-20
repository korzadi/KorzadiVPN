package wireguard

import (
	"crypto/rand"
	"encoding/base64"
)

func GenerateKey() string {
	key := make([]byte, 32)
	rand.Read(key)
	return base64.StdEncoding.EncodeToString(key)
}
