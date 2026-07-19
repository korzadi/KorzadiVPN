package wireguard

import "fmt"

type PeerController struct {
	Interface string
}

func NewPeerController() PeerController {

	return PeerController{
		Interface: "wg0",
	}
}

func (p PeerController) AddPeer(
	publicKey string,
	ip string,
) string {

	return fmt.Sprintf(
		"wg set %s peer %s allowed-ips %s/32",
		p.Interface,
		publicKey,
		ip,
	)
}

func (p PeerController) RemovePeer(
	publicKey string,
) string {

	return fmt.Sprintf(
		"wg set %s peer %s remove",
		p.Interface,
		publicKey,
	)
}
