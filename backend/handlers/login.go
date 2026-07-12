package handlers

import (
	"encoding/json"
	"net/http"
	"strings"
	"time"

	"golang.org/x/crypto/bcrypt"

	"korzadivpn/database"
	"korzadivpn/middleware"
	"korzadivpn/models"
)

func Login(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodPost {

		http.Error(
			w,
			"Metodo no permitido",
			http.StatusMethodNotAllowed,
		)

		return
	}

	var creds models.Credentials

	err := json.NewDecoder(
		r.Body,
	).Decode(&creds)

	if err != nil {

		http.Error(
			w,
			"JSON invalido",
			http.StatusBadRequest,
		)

		return
	}

	creds.Email = strings.TrimSpace(
		creds.Email,
	)

	user, err := database.GetUser(
		creds.Email,
	)

	if err != nil {

		http.Error(
			w,
			"Usuario o contraseña incorrectos",
			http.StatusUnauthorized,
		)

		return
	}

	err = bcrypt.CompareHashAndPassword(
		[]byte(user.Password),
		[]byte(creds.Password),
	)

	if err != nil {

		http.Error(
			w,
			"Usuario o contraseña incorrectos",
			http.StatusUnauthorized,
		)

		return
	}

	token, err := middleware.GenerateToken(
		user.Email,
	)

	if err != nil {

		http.Error(
			w,
			"Error creando token",
			http.StatusInternalServerError,
		)

		return
	}

	database.CreateActivity(
		models.Activity{

			Email: user.Email,

			Action: "login",

			IP: r.RemoteAddr,

			CreatedAt: time.Now().UTC().Format(time.RFC3339),
		},
	)

	w.Header().Set(
		"Content-Type",
		"application/json",
	)

	json.NewEncoder(w).Encode(
		map[string]interface{}{

			"message": "Login correcto",

			"token": token,

			"plan": user.Plan,

			"status": user.Status,
		},
	)

}
