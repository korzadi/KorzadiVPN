package handlers

import (
	"encoding/json"
	"net/http"
	"strings"
	"time"

	"golang.org/x/crypto/bcrypt"

	"korzadivpn/database"
	"korzadivpn/models"
)

func Register(w http.ResponseWriter, r *http.Request) {

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

	if creds.Email == "" || creds.Password == "" {

		http.Error(
			w,
			"Email y contraseña requeridos",
			http.StatusBadRequest,
		)

		return
	}

	hash, err := bcrypt.GenerateFromPassword(
		[]byte(creds.Password),
		bcrypt.DefaultCost,
	)

	if err != nil {

		http.Error(
			w,
			"Error creando contraseña",
			http.StatusInternalServerError,
		)

		return
	}

	user := models.User{

		Email: creds.Email,

		Password: string(hash),

		Plan: "free",

		Status: "active",
	}

	err = database.CreateUser(
		user,
	)

	if err != nil {

		http.Error(
			w,
			"Usuario ya existe",
			http.StatusConflict,
		)

		return
	}

	// Crear perfil VPN automáticamente

	err = CreateAutomaticVPNProfile(
		user.Email,
	)

	if err != nil {

		http.Error(
			w,
			"Usuario creado pero error creando VPN",
			http.StatusInternalServerError,
		)

		return
	}

	now := time.Now().UTC().Format(
		time.RFC3339,
	)

	database.CreateActivity(
		models.Activity{

			Email: user.Email,

			Action: "registered",

			CreatedAt: now,
		},
	)

	w.Header().Set(
		"Content-Type",
		"application/json",
	)

	json.NewEncoder(w).Encode(
		map[string]string{

			"message": "Usuario creado correctamente",

			"email": user.Email,

			"plan": user.Plan,

			"status": user.Status,

			"vpn": "created",
		},
	)
}
