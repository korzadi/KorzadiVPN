package handlers

import (
	"korzadivpn/database"
	"korzadivpn/models"
	"korzadivpn/vpncore/wireguard"
)

func CreateAutomaticVPNProfile(
	email string,
) error {

	privateKey := wireguard.GenerateKey()

	publicKey := wireguard.GenerateKey()

	client := wireguard.NewClient(
		email,
	)

	client.PrivateKey = privateKey

	client.PublicKey = publicKey

	client.Server = "vpn.korzadi.com"

	config := wireguard.GenerateConfig(
		client,
	)

	return database.CreateVPNProfile(
		models.VPNProfile{

			Email: email,

			ServerID: 1,

			Server: client.Server,

			Protocol: "wireguard",

			PublicKey: publicKey,

			PrivateKey: privateKey,

			Config: config,

			Status: "active",
		},
	)
}
