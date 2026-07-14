package handlers

import (
	"encoding/json"
	"net/http"

	"korzadivpn/database"
)

func AdminStats(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodGet {

		http.Error(
			w,
			"Método no permitido",
			http.StatusMethodNotAllowed,
		)

		return
	}

	stats, err := database.GetAdminStats()

	if err != nil {

		http.Error(
			w,
			"Error obteniendo estadísticas",
			http.StatusInternalServerError,
		)

		return
	}

	w.Header().Set(
		"Content-Type",
		"application/json",
	)

	json.NewEncoder(w).Encode(
		stats,
	)

}
