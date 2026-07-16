package wireguard

import "fmt"

func GenerateConfig(
	client Client,
) string {

	return fmt.Sprintf(
		`[Interface]
PrivateKey = %s
Address = %s
DNS = %s
MTU = 1420

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
