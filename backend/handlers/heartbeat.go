package handlers

import (
	"encoding/json"
	"net/http"

	"korzadivpn/database"
	"korzadivpn/middleware"
)

type HeartbeatRequest struct {
	ClientID string `json:"client_id"`
	Device   string `json:"device"`
}

func Heartbeat(w http.ResponseWriter, r *http.Request) {

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

	var req HeartbeatRequest

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "JSON inválido", http.StatusBadRequest)
		return
	}

	connection, err := database.GetConnectionByClientID(
		req.ClientID,
		email,
	)

	if err != nil {
		http.Error(w, "Error validando conexión", http.StatusInternalServerError)
		return
	}

	if connection == nil {
		http.Error(w, "Conexión no encontrada", http.StatusNotFound)
		return
	}

	if err := database.UpdateLastPing(req.ClientID); err != nil {
		http.Error(w, "Error actualizando heartbeat", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")

	json.NewEncoder(w).Encode(map[string]any{
		"status":    "ok",
		"client_id": req.ClientID,
		"device":    req.Device,
		"server":    connection.Server,
	})
}
