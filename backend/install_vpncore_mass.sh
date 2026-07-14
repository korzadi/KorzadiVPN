#!/data/data/com.termux/files/usr/bin/bash

set -e

echo "===== KORZADIVPN VPNCORE MASS INSTALL ====="

mkdir -p vpncore/wireguard

cat > vpncore/wireguard/client.go <<'EOC'
package wireguard

type Client struct {
	Email string
	Address string
	PublicKey string
	PrivateKey string
	Server string
	Port int
	DNS string
}

func NewClient(email string) Client {
	return Client{
		Email: email,
		Address: "10.8.0.2/32",
		Port: 51820,
		DNS: "1.1.1.1",
	}
}
EOC


cat > vpncore/wireguard/keys.go <<'EOC'
package wireguard

import (
	"crypto/rand"
	"encoding/base64"
)

func GenerateKey() string {
	key := make([]byte,32)
	rand.Read(key)
	return base64.StdEncoding.EncodeToString(key)
}
EOC


cat > vpncore/wireguard/config.go <<'EOC'
package wireguard

import "fmt"

func GenerateConfig(client Client) string {

	return fmt.Sprintf(
`[Interface]
PrivateKey = %s
Address = %s
DNS = %s

[Peer]
PublicKey = %s
Endpoint = %s:%d
AllowedIPs = 0.0.0.0/0, ::/0
PersistentKeepalive = 25
`,
	client.PrivateKey,
	client.Address,
	client.DNS,
	client.PublicKey,
	client.Server,
	client.Port,
	)
}
EOC


cat > vpncore/wireguard/profile.go <<'EOC'
package wireguard

func CreateProfile(email string, server string) Client {

	client := NewClient(email)

	client.Server = server
	client.PrivateKey = GenerateKey()
	client.PublicKey = GenerateKey()

	return client
}
EOC


gofmt -w vpncore/wireguard/*.go

echo "===== COMPILANDO ====="

go build ./...

echo "===== VPNCORE MASS INSTALADO ====="
