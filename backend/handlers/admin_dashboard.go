package handlers

import (
	"encoding/json"
	"net/http"

	"korzadivpn/database"
)

func AdminDashboard(
	w http.ResponseWriter,
	r *http.Request,
) {

	stats, err :=
		database.GetAdminStats()

	if err != nil {

		http.Error(
			w,
			"Error obteniendo estadísticas",
			http.StatusInternalServerError,
		)

		return
	}

	activity, _ :=
		database.GetActivity()

	w.Header().Set(
		"Content-Type",
		"application/json",
	)

	json.NewEncoder(w).Encode(
		map[string]interface{}{

			"statistics": stats,

			"activity": activity,
		},
	)

}
