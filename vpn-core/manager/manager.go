package manager

import (
	"fmt"
	"korzadivpn/vpn-core/wireguard"
)

func AddPeer(peer wireguard.Peer) error {

	fmt.Println(
		"Creando peer VPN:",
		peer.Email,
	)

	return nil
}

func RemovePeer(publicKey string) error {

	fmt.Println(
		"Eliminando peer:",
		publicKey,
	)

	return nil
}
