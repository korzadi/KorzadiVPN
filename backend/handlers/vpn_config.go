package handlers

import (
	"encoding/json"
	"net/http"

	"korzadivpn/database"
	"korzadivpn/middleware"
)

func GetVPNConfig(
	w http.ResponseWriter,
	r *http.Request,
) {

	if r.Method != http.MethodGet {

		http.Error(
			w,
			"Metodo no permitido",
			http.StatusMethodNotAllowed,
		)

		return
	}

	email, ok :=
		r.Context().
			Value(middleware.UserEmailKey).(string)

	if !ok {

		http.Error(
			w,
			"Usuario no autenticado",
			http.StatusUnauthorized,
		)

		return
	}

	client, err :=
		database.GetVPNClientByEmail(email)

	if err != nil {

		http.Error(
			w,
			"Cliente VPN no encontrado",
			http.StatusNotFound,
		)

		return
	}

	server, err :=
		database.GetServerByID(
			client.ServerID,
		)

	if err != nil {

		http.Error(
			w,
			"Servidor no encontrado",
			http.StatusNotFound,
		)

		return
	}

	json.NewEncoder(w).Encode(
		map[string]interface{}{

			"status": "ready",

			"client_ip": client.ClientIP,

			"client_public_key": client.PublicKey,

			"client_private_key": client.PrivateKey,

			"server": map[string]interface{}{

				"name": server.Name,

				"ip": server.ServerIP,

				"port": server.WireGuardPort,

				"public_key": server.ServerPublicKey,

				"dns": server.DNS,
			},
		},
	)

}
