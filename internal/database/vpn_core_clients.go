package database

import (
	"korzadivpn/internal/vpncore"
)

type VPNCoreClient struct {
	Email string

	IP string

	PublicKey string

	PrivateKey string
}

func CreateVPNCoreClient(
	client vpncore.ClientGenerator,
) error {

	return nil
}
