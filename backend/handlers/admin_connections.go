package handlers

import (
	"encoding/json"
	"net/http"

	"korzadivpn/database"
)

// AdminConnections muestra conexiones activas.
func AdminConnections(
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

	connections, err := database.GetAllActiveConnections()

	if err != nil {

		http.Error(
			w,
			err.Error(),
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
			"total":       len(connections),
			"connections": connections,
		},
	)
}
