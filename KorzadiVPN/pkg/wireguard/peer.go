package wireguard

import "fmt"

type Peer struct {
	PublicKey string

	AllowedIP string

	ClientIP string

	Description string
}

func NewPeer(
	publicKey string,
	clientIP string,
) Peer {

	return Peer{

		PublicKey: publicKey,

		AllowedIP: clientIP + "/32",

		ClientIP: clientIP,

		Description: "KorzadiVPN Client",
	}

}

func ValidatePeer(
	peer Peer,
) error {

	if peer.PublicKey == "" {

		return fmt.Errorf(
			"public key vacía",
		)

	}

	if peer.ClientIP == "" {

		return fmt.Errorf(
			"client ip vacía",
		)

	}

	return nil

}
