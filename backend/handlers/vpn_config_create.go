package handlers

import (
	"encoding/json"
	"net/http"

	"korzadivpn/database"
	"korzadivpn/middleware"
	"korzadivpn/vpncore/wireguard"
)

func CreateVPNConfig(
	w http.ResponseWriter,
	r *http.Request,
) {

	email, ok := r.Context().
		Value(middleware.UserEmailKey).(string)

	if !ok {

		http.Error(
			w,
			"Usuario no autenticado",
			http.StatusUnauthorized,
		)

		return
	}

	client, err := database.GetVPNClientByEmail(
		email,
	)

	if err != nil {

		http.Error(
			w,
			"Cliente VPN no encontrado",
			http.StatusNotFound,
		)

		return
	}

	server, err := database.GetServerByID(
		client.ServerID,
	)

	if err != nil {

		http.Error(
			w,
			"Servidor VPN no encontrado",
			http.StatusInternalServerError,
		)

		return
	}

	wgClient := wireguard.Client{

		Email: email,

		Address: client.ClientIP + "/32",

		PrivateKey: client.PrivateKey,

		PublicKey: server.ServerPublicKey,

		Server: server.ServerIP,

		Port: server.WireGuardPort,

		DNS: client.DNS,
	}

	config := wireguard.GenerateConfig(
		wgClient,
	)

	err = database.UpdateVPNClientConfig(
		client.ID,
		config,
	)

	if err != nil {

		http.Error(
			w,
			"Error guardando configuración VPN",
			http.StatusInternalServerError,
		)

		return
	}

	w.Header().Set(
		"Content-Type",
		"application/json",
	)

	json.NewEncoder(w).Encode(
		map[string]interface{}{

			"status": "created",

			"email": email,

			"config": config,
		},
	)

}
