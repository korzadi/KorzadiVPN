package handlers

import (
	"encoding/json"
	"net/http"
	"time"

	"korzadivpn/internal/database"
	"korzadivpn/internal/models"
)

type AdminChangePlanRequest struct {
	Email string `json:"email"`

	Plan string `json:"plan"`
}

func AdminChangePlan(
	w http.ResponseWriter,
	r *http.Request,
) {

	if r.Method != http.MethodPost {

		http.Error(
			w,
			"Metodo no permitido",
			http.StatusMethodNotAllowed,
		)

		return
	}

	var req AdminChangePlanRequest

	err := json.NewDecoder(
		r.Body,
	).Decode(&req)

	if err != nil {

		http.Error(
			w,
			"JSON invalido",
			http.StatusBadRequest,
		)

		return
	}

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
			"Plan no valido",
			http.StatusBadRequest,
		)

		return
	}

	err = database.UpdateUserPlan(
		req.Email,
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

	database.CreateActivity(
		models.Activity{

			Email: req.Email,

			Action: "admin_change_plan",

			IP: r.RemoteAddr,

			CreatedAt: time.Now().
				UTC().
				Format(time.RFC3339),
		},
	)

	user, err :=
		database.GetUser(req.Email)

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

	json.NewEncoder(w).Encode(
		user,
	)

}
