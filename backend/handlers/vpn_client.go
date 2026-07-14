package handlers

import (
	"encoding/json"
	"net/http"
	"time"

	"korzadivpn/database"
	"korzadivpn/middleware"
	"korzadivpn/models"
	"korzadivpn/utils"
)

func CreateVPNClient(
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

	server, err :=
		database.GetBestServer()

	if err != nil {

		http.Error(
			w,
			"No hay servidores disponibles",
			http.StatusInternalServerError,
		)

		return
	}

	clientIP, err :=
		database.GetNextVPNClientIP()

	if err != nil {

		http.Error(
			w,
			"No hay IP disponible",
			http.StatusInternalServerError,
		)

		return
	}

	publicKey, privateKey :=
		utils.GenerateWireGuardKeys()

	now :=
		time.Now().
			UTC().
			Format(time.RFC3339)

	client := models.VPNClient{

		Email: email,

		ServerID: server.ID,

		NodeID: server.ID,

		ClientName: "Korzadi-Device",

		DeviceName: "Korzadi Device",

		DeviceType: "WireGuard",

		ClientIP: clientIP,

		PublicKey: publicKey,

		PrivateKey: privateKey,

		Protocol: "wireguard",

		DNS: "1.1.1.1",

		MTU: 1420,

		AllowedIPs: "0.0.0.0/0, ::/0",

		Endpoint: server.Name,

		Status: "active",

		ConnectionStatus: "offline",

		Plan: "free",

		BandwidthLimit: 0,

		DataUsed: 0,

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
			"Error creando cliente VPN",
			http.StatusInternalServerError,
		)

		return
	}

	database.IncrementServerUsers(
		server.ID,
	)

	device := models.Device{

		Email: email,

		DeviceName: client.DeviceName,

		DeviceType: client.DeviceType,

		Status: "active",

		LastServer: server.Name,

		LastSeen: now,

		CreatedAt: now,
	}

	err =
		database.UpsertDevice(
			device,
		)

	if err != nil {

		http.Error(
			w,
			"Error registrando dispositivo",
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

			"message": "Cliente VPN creado correctamente",

			"client": client,

			"server": server,

			"device": device,
		},
	)

}
