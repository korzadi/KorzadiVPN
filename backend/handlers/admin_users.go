package handlers

import (
	"encoding/json"
	"net/http"

	"korzadivpn/database"
)

// AdminUsers muestra usuarios del sistema.
func AdminUsers(
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

	users, err :=
		database.GetUsers()

	if err != nil {

		http.Error(
			w,
			"Error obteniendo usuarios",
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

			"total": len(users),

			"users": users,
		},
	)

}
