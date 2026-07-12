package handlers

import (
	"encoding/json"
	"net/http"

	"korzadivpn/database"
	"korzadivpn/middleware"
)

func UserDevices(w http.ResponseWriter, r *http.Request) {

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

	devices, err := database.GetDevicesByEmail(email)

	if err != nil {
		http.Error(
			w,
			"Error obteniendo dispositivos",
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
			"devices": devices,
			"total":   len(devices),
		},
	)
}
