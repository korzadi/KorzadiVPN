package wireguard

import "fmt"

func AddPeerCommand(
	interfaceName string,
	peer Peer,
) string {

	return fmt.Sprintf(
		"wg set %s peer %s allowed-ips %s",
		interfaceName,
		peer.PublicKey,
		peer.AllowedIP,
	)

}

func RemovePeerCommand(
	interfaceName string,
	publicKey string,
) string {

	return fmt.Sprintf(
		"wg set %s peer %s remove",
		interfaceName,
		publicKey,
	)

}
