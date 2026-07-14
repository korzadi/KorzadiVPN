package handlers

import (
	"encoding/json"
	"net/http"

	"korzadivpn/vpncore/connection"
)

func StartVPNConnection(
	w http.ResponseWriter,
	r *http.Request,
) {

	var data struct {
		Email string `json:"email"`

		Server string `json:"server"`

		IP string `json:"ip"`
	}

	json.NewDecoder(
		r.Body,
	).Decode(&data)

	conn := connection.Start(
		data.Email,
		data.Server,
		data.IP,
	)

	w.Header().Set(
		"Content-Type",
		"application/json",
	)

	json.NewEncoder(w).Encode(conn)
}
