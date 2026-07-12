package handlers

import (
	"encoding/json"
	"net/http"

	"korzadivpn/database"
)

// AdminServers muestra todos los servidores para administración.
func AdminServers(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodGet {

		http.Error(
			w,
			"Método no permitido",
			http.StatusMethodNotAllowed,
		)

		return
	}

	servers, err := database.GetServers()

	if err != nil {

		http.Error(
			w,
			"Error obteniendo servidores",
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

			"total": len(servers),

			"servers": servers,
		},
	)

}
