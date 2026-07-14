package handlers

import (
	"encoding/json"
	"net/http"

	"korzadivpn/models"
)

func AppConfig(
	w http.ResponseWriter,
	r *http.Request,
) {

	w.Header().Set(
		"Content-Type",
		"application/json",
	)

	json.NewEncoder(w).Encode(
		models.AppConfig{

			Version: "1.0",

			API: "KorzadiVPN",

			VPNCore: "WireGuard",

			Status: "ready",
		},
	)
}
