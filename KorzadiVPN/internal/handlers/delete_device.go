package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"
	"strings"
	"time"

	"korzadivpn/internal/database"
	"korzadivpn/internal/middleware"
	"korzadivpn/internal/models"
)

func DeleteDevice(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodDelete {

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
		"/api/user/device/",
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

	devices, err := database.GetDevicesByEmail(email)

	if err != nil {

		http.Error(
			w,
			"Error buscando dispositivos",
			http.StatusInternalServerError,
		)

		return
	}

	var device *models.Device

	for i := range devices {

		if devices[i].ID == id {

			device = &devices[i]
			break
		}
	}

	if device == nil {

		http.Error(
			w,
			"Dispositivo no encontrado",
			http.StatusNotFound,
		)

		return
	}

	err = database.DeleteDevice(
		id,
		email,
	)

	if err != nil {

		http.Error(
			w,
			"Error eliminando dispositivo",
			http.StatusInternalServerError,
		)

		return
	}

	database.CreateActivity(
		models.Activity{

			Email: email,

			Server: device.LastServer,

			Action: "device_deleted",

			Device: device.DeviceName,

			IP: device.LastIP,

			CreatedAt: time.Now().
				UTC().
				Format(time.RFC3339),
		},
	)

	w.Header().
		Set(
			"Content-Type",
			"application/json",
		)

	json.NewEncoder(w).Encode(
		map[string]interface{}{

			"message": "Dispositivo eliminado correctamente",

			"device": device.DeviceName,
		},
	)

}
