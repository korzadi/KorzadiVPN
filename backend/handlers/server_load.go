package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"korzadivpn/database"
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

	servers, err := database.GetServers()

	if err != nil {

		http.Error(
			w,
			"Error obteniendo servidores",
			http.StatusInternalServerError,
		)

		return
	}

	type ServerLoad struct {
		ID int `json:"id"`

		Name string `json:"name"`

		Country string `json:"country"`

		City string `json:"city"`

		Status string `json:"status"`

		Protocol string `json:"protocol"`

		CurrentUsers int `json:"current_users"`

		MaxUsers int `json:"max_users"`

		Load string `json:"load"`

		Latency int `json:"latency"`
	}

	result := make([]ServerLoad, 0)

	for _, server := range servers {

		load := 0

		if server.MaxUsers > 0 {

			load = (server.CurrentUsers * 100) / server.MaxUsers

		}

		result = append(
			result,
			ServerLoad{

				ID: server.ID,

				Name: server.Name,

				Country: server.Country,

				City: server.City,

				Status: server.Status,

				Protocol: server.Protocol,

				CurrentUsers: server.CurrentUsers,

				MaxUsers: server.MaxUsers,

				Load: strconv.Itoa(load) + "%",

				Latency: server.Latency,
			},
		)

	}

	w.Header().Set(
		"Content-Type",
		"application/json",
	)

	json.NewEncoder(w).Encode(
		result,
	)

}
