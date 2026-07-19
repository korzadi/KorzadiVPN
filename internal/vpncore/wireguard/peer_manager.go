package wireguard

import (
	"fmt"

	"korzadivpn/internal/models"
)

type PeerConfig struct {
	PublicKey string

	AllowedIP string
}

func CreatePeerFromClient(
	client models.VPNClient,
) PeerConfig {

	return PeerConfig{

		PublicKey: client.PublicKey,

		AllowedIP: client.ClientIP,
	}

}

func GenerateAddPeerCommand(
	interfaceName string,
	peer PeerConfig,
) string {

	return fmt.Sprintf(
		"wg set %s peer %s allowed-ips %s/32",
		interfaceName,
		peer.PublicKey,
		peer.AllowedIP,
	)

}
