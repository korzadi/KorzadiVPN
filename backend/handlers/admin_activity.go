package handlers

import (
	"encoding/json"
	"net/http"

	"korzadivpn/database"
)

// AdminActivity muestra historial de actividad.
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

	w.Header().Set(
		"Content-Type",
		"application/json",
	)

	json.NewEncoder(w).Encode(
		map[string]interface{}{

			"total": len(activities),

			"activities": activities,
		},
	)
}
