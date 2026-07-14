package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"
	"strings"
	"time"

	"korzadivpn/database"
	"korzadivpn/middleware"
	"korzadivpn/models"
)

func DisconnectDevice(w http.ResponseWriter, r *http.Request) {

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

	idText := strings.TrimPrefix(
		r.URL.Path,
		"/api/user/device/disconnect/",
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

	err = database.DisconnectDeviceConnection(
		email,
		device.DeviceName,
	)

	if err != nil {

		http.Error(
			w,
			"Error desconectando dispositivo",
			http.StatusInternalServerError,
		)

		return
	}

	device.Status = "offline"
	device.LastSeen = time.Now().
		UTC().
		Format(time.RFC3339)

	database.UpdateDevice(
		*device,
	)

	database.CreateActivity(
		models.Activity{

			Email: email,

			Server: device.LastServer,

			Action: "device_disconnected",

			Device: device.DeviceName,

			IP: device.LastIP,

			CreatedAt: time.Now().
				UTC().
				Format(time.RFC3339),
		},
	)

	w.Header().Set(
		"Content-Type",
		"application/json",
	)

	json.NewEncoder(w).Encode(
		map[string]interface{}{

			"message": "Dispositivo desconectado correctamente",

			"device": device.DeviceName,
		},
	)

}
