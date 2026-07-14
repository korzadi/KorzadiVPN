package main

import (
	"fmt"
	"log"

	"korzadivpn/database"
	"korzadivpn/models"
	"korzadivpn/vpncore/wireguard"
)

func main() {

	database.Connect()

	total := 100

	for i := 1; i <= total; i++ {

		email := fmt.Sprintf(
			"bulk%d@korzadi.com",
			i,
		)

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

		err := database.CreateVPNProfile(
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

		if err != nil {
			log.Println("ERROR:", email, err)
			continue
		}

		fmt.Println(
			"CREADO:",
			email,
		)
	}

	fmt.Println(
		"CREACION MASIVA TERMINADA",
	)
}
