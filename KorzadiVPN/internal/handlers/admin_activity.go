package handlers

import (
	"encoding/json"
	"net/http"

	"korzadivpn/internal/database"
)

// AdminActivity muestra historial y estadísticas de actividad.
func AdminActivity(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodGet {

		http.Error(
			w,
			"Método no permitido",
			http.StatusMethodNotAllowed,
		)

		return
	}

	activities, err := database.GetActivity()

	if err != nil {

		http.Error(
			w,
			"Error obteniendo actividad",
			http.StatusInternalServerError,
		)

		return
	}

	total, _ := database.CountActivities()

	connects, _ := database.CountActivityByAction(
		"connected",
	)

	disconnects, _ := database.CountActivityByAction(
		"disconnected",
	)

	lastActivity, _ := database.GetLastActivity()

	w.Header().Set(
		"Content-Type",
		"application/json",
	)

	json.NewEncoder(w).Encode(
		map[string]interface{}{

			"statistics": map[string]interface{}{

				"total_events": total,

				"total_connections": connects,

				"total_disconnects": disconnects,

				"last_activity": lastActivity,
			},

			"total": len(activities),

			"activities": activities,
		},
	)

}
