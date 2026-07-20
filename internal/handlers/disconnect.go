package handlers

import (
	"encoding/json"
	"net/http"
	"time"

	"korzadivpn/internal/database"
	"korzadivpn/internal/middleware"
	"korzadivpn/internal/models"
)

// Disconnect desconecta al usuario actual.
func Disconnect(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodPost {

		http.Error(
			w,
			"Método no permitido",
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

	connection, err := database.GetActiveConnection(email)

	if err != nil {

		http.Error(
			w,
			"Error buscando conexión: "+err.Error(),
			http.StatusInternalServerError,
		)

		return
	}

	if connection == nil {

		http.Error(
			w,
			"No existe conexión activa",
			http.StatusNotFound,
		)

		return
	}

	err = database.DisconnectConnection(email)

	if err != nil {

		http.Error(
			w,
			"Error desconectando usuario",
			http.StatusInternalServerError,
		)

		return
	}

	device, err := database.GetDeviceByName(
		email,
		connection.Device,
	)

	if err == nil && device != nil {

		device.Status = "offline"
		device.LastSeen = time.Now().UTC().Format(time.RFC3339)

		database.UpdateDevice(*device)
	}

	database.CreateActivity(
		models.Activity{

			Email: email,

			Server: connection.Server,

			Action: "disconnected",

			Device: connection.Device,

			IP: connection.IP,

			CreatedAt: time.Now().UTC().Format(time.RFC3339),
		},
	)

	database.DecrementServerUsers(
		connection.ServerID,
	)

	w.Header().Set(
		"Content-Type",
		"application/json",
	)

	json.NewEncoder(w).Encode(
		map[string]interface{}{

			"message": "Desconectado correctamente",

			"server": connection.Server,

			"device": connection.Device,

			"disconnected_at": time.Now().UTC().Format(time.RFC3339),
		},
	)

}
