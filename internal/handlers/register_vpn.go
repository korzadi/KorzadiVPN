package handlers

import (
	"korzadivpn/internal/database"
	"korzadivpn/internal/models"
	"korzadivpn/internal/vpncore/wireguard"
	"korzadivpn/pkg/utils"
)

func CreateAutomaticVPNProfile(
	email string,
) error {

	publicKey, privateKey :=
		utils.GenerateWireGuardKeys()

	client := wireguard.NewClient(
		email,
	)

	client.PrivateKey = privateKey

	client.PublicKey = publicKey

	client.Server = "pending-vps"

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

			Status: "ready",
		},
	)
}
