package handlers

import (
	"encoding/json"
	"net/http"

	"korzadivpn/internal/database"
	"korzadivpn/internal/middleware"
	"korzadivpn/internal/services"
)

func StartVPNConnection(
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

	client, err :=
		database.GetVPNClientByEmail(
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

	service :=
		services.NewVPNService()

	err =
		service.Connect(
			client,
		)

	if err != nil {

		http.Error(
			w,
			"Error conectando VPN: "+err.Error(),
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

			"message": "VPN conectada correctamente",

			"email": email,

			"status": "connected",
		},
	)

}
