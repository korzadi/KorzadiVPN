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

type ConnectRequest struct {
	ServerID int    `json:"server_id"`
	Device   string `json:"device"`
}

func Connect(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodPost {
		http.Error(w, "Método no permitido", http.StatusMethodNotAllowed)
		return
	}

	email, ok := r.Context().Value(
		middleware.UserEmailKey,
	).(string)

	if !ok {
		http.Error(w, "Usuario no autenticado", http.StatusUnauthorized)
		return
	}

	user, err := database.GetUser(email)

	if err != nil {
		http.Error(w, "Usuario no encontrado", http.StatusNotFound)
		return
	}

	var req ConnectRequest

	err = json.NewDecoder(r.Body).Decode(&req)

	if err != nil {
		http.Error(w, "JSON inválido", http.StatusBadRequest)
		return
	}

	limit := models.GetDeviceLimit(user.Plan)

	active, err := database.CountActiveDevices(email)

	if err != nil {
		http.Error(w, "Error consultando dispositivos", http.StatusInternalServerError)
		return
	}

	if active >= limit {

		existing, _ := database.GetDeviceByName(
			email,
			req.Device,
		)

		if existing == nil {

			http.Error(
				w,
				"Límite de dispositivos alcanzado",
				http.StatusForbidden,
			)

			return
		}
	}

	var server *models.Server

	if req.ServerID == 0 {

		server, err = database.GetBestServer()

		if err != nil {
			http.Error(
				w,
				"No hay servidores disponibles",
				http.StatusServiceUnavailable,
			)

			return
		}

	} else {

		server, err = database.GetServerByID(req.ServerID)

		if err != nil {

			http.Error(
				w,
				"Servidor no encontrado",
				http.StatusNotFound,
			)

			return
		}
	}

	now := time.Now().UTC().Format(time.RFC3339)

	err = database.UpsertDevice(
		models.Device{
			Email:      email,
			DeviceName: req.Device,
			DeviceType: req.Device,
			Status:     "connected",
			LastIP:     r.RemoteAddr,
			LastServer: server.Name,
			LastSeen:   now,
			CreatedAt:  now,
		},
	)

	if err != nil {
		http.Error(
			w,
			"Error guardando dispositivo",
			http.StatusInternalServerError,
		)
		return
	}

	connection := models.Connection{

		Email: email,

		ServerID: server.ID,

		Server: server.Name,

		Status: "connected",

		Device: req.Device,

		ClientID: utils.GenerateClientID(),

		IP: r.RemoteAddr,

		ConnectedAt: now,

		LastPing: now,
	}

	err = database.CreateConnection(connection)

	if err != nil {
		http.Error(
			w,
			"Error creando conexión",
			http.StatusInternalServerError,
		)
		return
	}

	database.CreateActivity(
		models.Activity{
			Email:     email,
			Server:    server.Name,
			Action:    "connected",
			Device:    req.Device,
			IP:        r.RemoteAddr,
			CreatedAt: now,
		},
	)

	database.IncrementServerUsers(
		server.ID,
	)

	w.Header().Set(
		"Content-Type",
		"application/json",
	)

	json.NewEncoder(w).Encode(
		map[string]interface{}{
			"message":      "Conectado correctamente",
			"server":       server.Name,
			"device":       req.Device,
			"connected_at": now,
		},
	)
}
