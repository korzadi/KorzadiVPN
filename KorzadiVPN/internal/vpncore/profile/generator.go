package profile

import (
	"fmt"
)

func GenerateWireGuardConfig(
	privateKey string,
	address string,
	dns string,
	publicKey string,
	endpoint string,
) string {

	return fmt.Sprintf(
		`[Interface]
PrivateKey = %s
Address = %s
DNS = %s

[Peer]
PublicKey = %s
Endpoint = %s
AllowedIPs = 0.0.0.0/0, ::/0
PersistentKeepalive = 25
`,
		privateKey,
		address,
		dns,
		publicKey,
		endpoint,
	)
}
