package generator

import (
	"fmt"

	"korzadivpn/vpn-core/wireguard"
)

func GenerateClientConfig(
	peer wireguard.Peer,
	server wireguard.Server,
) string {

	config := fmt.Sprintf(
		`[Interface]
PrivateKey = %s
Address = %s

[Peer]
PublicKey = %s
Endpoint = %s:%d
AllowedIPs = 0.0.0.0/0
PersistentKeepalive = 25
`,
		peer.PrivateKey,
		peer.IPAddress,
		server.PublicKey,
		server.Endpoint,
		server.Port,
	)

	return config
}
