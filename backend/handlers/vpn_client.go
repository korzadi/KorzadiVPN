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

func CreateVPNClient(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodPost {
		http.Error(w, "Método no permitido", http.StatusMethodNotAllowed)
		return
	}

	email, ok := r.Context().
		Value(middleware.UserEmailKey).(string)

	if !ok {
		http.Error(w, "Usuario no autenticado", http.StatusUnauthorized)
		return
	}

	publicKey, privateKey := utils.GenerateWireGuardKeys()

	client := models.VPNClient{

		Email: email,

		ServerID: 1,

		ClientName: "Korzadi-Device",

		ClientIP: "10.0.0.2",

		PublicKey: publicKey,

		PrivateKey: privateKey,

		Status: "active",

		CreatedAt: time.Now().UTC().Format(time.RFC3339),
	}


	err := database.CreateVPNClient(client)

	if err != nil {

		http.Error(
			w,
			"Error creando cliente VPN",
			http.StatusInternalServerError,
		)

		return
	}


	w.Header().Set(
		"Content-Type",
		"application/json",
	)


	json.NewEncoder(w).Encode(
		client,
	)

}
