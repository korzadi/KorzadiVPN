package handlers

import (
	"encoding/json"
	"net/http"

	"korzadivpn/internal/middleware"
)

func Profile(w http.ResponseWriter, r *http.Request) {

	email := r.Context().Value(
		middleware.UserEmailKey,
	).(string)

	json.NewEncoder(w).Encode(map[string]string{
		"email":  email,
		"plan":   "free",
		"status": "active",
	})
}
