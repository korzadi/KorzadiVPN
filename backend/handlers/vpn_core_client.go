package handlers

import (
	"encoding/json"
	"net/http"

	"korzadivpn/database"
	"korzadivpn/middleware"
	"korzadivpn/vpncore"
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

	client := vpncore.CreateClient(
		email,
	)

	err := database.CreateVPNCoreClient(
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

	w.Header().Set(
		"Content-Type",
		"application/json",
	)

	json.NewEncoder(w).Encode(
		map[string]interface{}{

			"status": "created",

			"client": client,
		},
	)
}
