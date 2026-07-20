package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"
	"strings"

	"korzadivpn/internal/database"
	"korzadivpn/internal/middleware"
)

func DeviceDetail(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodGet {

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

	idText := strings.TrimPrefix(
		r.URL.Path,
		"/api/user/device/view/",
	)

	id, err := strconv.Atoi(idText)

	if err != nil {

		http.Error(
			w,
			"ID invalido",
			http.StatusBadRequest,
		)

		return
	}

	device, err := database.GetDeviceByID(
		id,
		email,
	)

	if err != nil || device == nil {

		http.Error(
			w,
			"Dispositivo no encontrado",
			http.StatusNotFound,
		)

		return
	}

	w.Header().Set(
		"Content-Type",
		"application/json",
	)

	json.NewEncoder(w).Encode(
		device,
	)

}
