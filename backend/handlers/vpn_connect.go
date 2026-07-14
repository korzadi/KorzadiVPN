package handlers

import (
	"encoding/json"
	"net/http"
	"time"

	"korzadivpn/database"
	"korzadivpn/middleware"
)

func VPNConnect(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodPost {
		http.Error(
			w,
			"Metodo no permitido",
			http.StatusMethodNotAllowed,
		)
		return
	}

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

	client, err := database.GetVPNClientByEmail(email)

	if err != nil {

		http.Error(
			w,
			"No existe cliente VPN",
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
			"Servidor no encontrado",
			http.StatusNotFound,
		)

		return
	}

	now :=
		time.Now().
			UTC().
			Format(time.RFC3339)

	w.Header().
		Set(
			"Content-Type",
			"application/json",
		)

	json.NewEncoder(w).Encode(
		map[string]interface{}{

			"status": "connected",

			"server": server.Name,

			"protocol": server.Protocol,

			"client_ip": client.ClientIP,

			"server_ip": server.ServerIP,

			"port": server.WireGuardPort,

			"dns": server.DNS,

			"connected_at": now,
		},
	)

}
