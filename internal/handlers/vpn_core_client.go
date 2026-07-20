package handlers

import (
	"encoding/json"
	"net/http"
	"time"

	"korzadivpn/internal/database"
	"korzadivpn/internal/middleware"
	"korzadivpn/internal/models"
	"korzadivpn/internal/vpncore"
)

func CreateVPNCoreClient(
	w http.ResponseWriter,
	r *http.Request,
) {

	if r.Method != http.MethodPost {

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

	ip, err :=
		database.GetNextVPNClientIP()

	if err != nil {

		http.Error(
			w,
			"IP no disponible",
			http.StatusInternalServerError,
		)

		return
	}

	generated :=
		vpncore.CreateClient(
			email,
			ip,
		)

	server, err :=
		database.GetBestServer()

	if err != nil {

		http.Error(
			w,
			"Servidor no disponible",
			http.StatusInternalServerError,
		)

		return
	}

	now :=
		time.Now().
			UTC().
			Format(time.RFC3339)

	client :=
		models.VPNClient{

			Email: email,

			ServerID: server.ID,

			NodeID: server.ID,

			ClientName: "Korzadi-Core",

			DeviceName: "Korzadi Core Device",

			DeviceType: "WireGuard",

			ClientIP: generated.IP,

			PublicKey: generated.PublicKey,

			PrivateKey: generated.PrivateKey,

			Protocol: "wireguard",

			DNS: "1.1.1.1",

			MTU: 1420,

			AllowedIPs: "0.0.0.0/0, ::/0",

			Endpoint: server.ServerIP,

			Status: "active",

			ConnectionStatus: "offline",

			Plan: "free",

			MaxDevices: 1,

			CreatedAt: now,

			UpdatedAt: now,
		}

	err =
		database.CreateVPNClient(
			client,
		)

	if err != nil {

		http.Error(
			w,
			"Error creando cliente",
			http.StatusInternalServerError,
		)

		return
	}

	json.NewEncoder(w).
		Encode(
			map[string]interface{}{

				"status": "created",

				"client": client,
			},
		)
}
