package handlers

import (
	"encoding/json"
	"net/http"
	"strings"
	"time"

	"korzadivpn/database"
	"korzadivpn/middleware"
	"korzadivpn/models"
)

func Logout(
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

	email, ok := r.Context().
		Value(middleware.UserEmailKey).(string)

	if !ok {

		http.Error(
			w,
			"Usuario no autenticado",
			http.StatusUnauthorized,
		)

		return
	}

	authHeader := r.Header.Get(
		"Authorization",
	)

	token := strings.TrimPrefix(
		authHeader,
		"Bearer ",
	)

	if token == "" {

		http.Error(
			w,
			"Token requerido",
			http.StatusUnauthorized,
		)

		return
	}

	err := database.RevokeSession(
		token,
	)

	if err != nil {

		http.Error(
			w,
			"Error cerrando sesión",
			http.StatusInternalServerError,
		)

		return
	}

	database.CreateActivity(
		models.Activity{

			Email: email,

			Action: "logout",

			IP: r.RemoteAddr,

			CreatedAt: time.Now().
				UTC().
				Format(time.RFC3339),
		},
	)

	w.Header().Set(
		"Content-Type",
		"application/json",
	)

	json.NewEncoder(w).Encode(
		map[string]string{

			"message": "Sesion cerrada correctamente",
		},
	)

}
