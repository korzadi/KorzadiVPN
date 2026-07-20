package handlers

import (
	"encoding/json"
	"net/http"
)

func VPNCoreStatus(
	w http.ResponseWriter,
	r *http.Request,
) {

	w.Header().Set(
		"Content-Type",
		"application/json",
	)

	json.NewEncoder(w).Encode(
		map[string]interface{}{
			"status": "online",
			"core":   "wireguard",
			"nodes":  "ready",
		},
	)
}
