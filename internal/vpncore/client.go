package vpncore

import (
	"korzadivpn/pkg/utils"
)

type ClientGenerator struct {
	Email string

	IP string

	PublicKey string

	PrivateKey string
}

func CreateClient(
	email string,
	ip string,
) ClientGenerator {

	publicKey, privateKey :=
		utils.GenerateWireGuardKeys()

	return ClientGenerator{

		Email: email,

		IP: ip,

		PublicKey: publicKey,

		PrivateKey: privateKey,
	}
}
