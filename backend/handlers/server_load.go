package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"korzadivpn/models"
)

func ServersLoad(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodGet {
		http.Error(
			w,
			"Método no permitido",
			http.StatusMethodNotAllowed,
		)
		return
	}

	type ServerLoad struct {
		Name         string `json:"name"`
		CurrentUsers int    `json:"current_users"`
		MaxUsers     int    `json:"max_users"`
		Load         string `json:"load"`
	}

	var result []ServerLoad

	for _, server := range models.Servers {

		load := 0

		if server.MaxUsers > 0 {

			load = (server.CurrentUsers * 100) / server.MaxUsers
		}

		result = append(
			result,
			ServerLoad{

				Name: server.Name,

				CurrentUsers: server.CurrentUsers,

				MaxUsers: server.MaxUsers,

				Load: strconv.Itoa(load) + "%",
			},
		)
	}

	w.Header().Set(
		"Content-Type",
		"application/json",
	)

	json.NewEncoder(w).Encode(result)
}
