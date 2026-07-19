package handlers

import (
	"encoding/json"
	"net/http"

	"korzadivpn/internal/database"
	"korzadivpn/internal/middleware"
)

func VPNStatus(
	w http.ResponseWriter,
	r *http.Request,
) {

	if r.Method != http.MethodGet {

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

	connection, err :=
		database.GetActiveConnection(email)

	if err != nil {

		http.Error(
			w,
			"Error buscando conexión",
			http.StatusInternalServerError,
		)

		return
	}

	if connection == nil {

		w.Header().Set(
			"Content-Type",
			"application/json",
		)

		json.NewEncoder(w).Encode(
			map[string]interface{}{
				"status": "disconnected",
			},
		)

		return
	}

	w.Header().Set(
		"Content-Type",
		"application/json",
	)

	json.NewEncoder(w).Encode(
		map[string]interface{}{

			"status": connection.Status,

			"server": connection.Server,

			"device": connection.Device,

			"client_id": connection.ClientID,

			"connected_at": connection.ConnectedAt,

			"last_ping": connection.LastPing,
		},
	)

}
