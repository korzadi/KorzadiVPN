package handlers

import (
	"encoding/json"
	"net/http"

	"korzadivpn/database"
	"korzadivpn/middleware"
	"korzadivpn/models"
)

type ChangePlanRequest struct {
	Plan string `json:"plan"`
}

func ChangePlan(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodPost {

		http.Error(
			w,
			"Método no permitido",
			http.StatusMethodNotAllowed,
		)

		return
	}

	email, ok := r.Context().Value(
		middleware.UserEmailKey,
	).(string)

	if !ok {

		http.Error(
			w,
			"Usuario no autenticado",
			http.StatusUnauthorized,
		)

		return
	}

	var req ChangePlanRequest

	err := json.NewDecoder(
		r.Body,
	).Decode(&req)

	if err != nil {

		http.Error(
			w,
			"JSON inválido",
			http.StatusBadRequest,
		)

		return
	}

	// Validar plan

	valid := false

	for _, plan := range models.Plans {

		if plan.Name == req.Plan {

			valid = true

			break
		}

	}

	if !valid {

		http.Error(
			w,
			"Plan no válido",
			http.StatusBadRequest,
		)

		return
	}

	// Actualizar en SQLite

	err = database.UpdateUserPlan(
		email,
		req.Plan,
	)

	if err != nil {

		http.Error(
			w,
			"Error actualizando plan",
			http.StatusInternalServerError,
		)

		return
	}

	user, err := database.GetUser(email)

	if err != nil {

		http.Error(
			w,
			"Usuario no encontrado",
			http.StatusNotFound,
		)

		return
	}

	w.Header().Set(
		"Content-Type",
		"application/json",
	)

	json.NewEncoder(w).Encode(user)

}
