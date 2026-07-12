package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"
	"strings"

	"korzadivpn/database"
	"korzadivpn/models"
)

// Lista todos los servidores
func Servers(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodGet {
		http.Error(w, "Método no permitido", http.StatusMethodNotAllowed)
		return
	}

	servers, err := database.GetServers()

	if err != nil {
		http.Error(w, "Error obteniendo servidores", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")

	json.NewEncoder(w).Encode(servers)
}

// Servidor por ID
func Server(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodGet {
		http.Error(w, "Método no permitido", http.StatusMethodNotAllowed)
		return
	}

	idStr := r.URL.Query().Get("id")

	if idStr == "" {
		http.Error(w, "Falta el parámetro id", http.StatusBadRequest)
		return
	}

	id, err := strconv.Atoi(idStr)

	if err != nil {
		http.Error(w, "ID inválido", http.StatusBadRequest)
		return
	}

	servers, err := database.GetServers()

	if err != nil {
		http.Error(w, "Error obteniendo servidores", http.StatusInternalServerError)
		return
	}

	for _, server := range servers {

		if server.ID == id {

			w.Header().Set("Content-Type", "application/json")

			json.NewEncoder(w).Encode(server)

			return
		}

	}

	http.Error(
		w,
		"Servidor no encontrado",
		http.StatusNotFound,
	)

}

// Buscar por país
func ServersByCountry(w http.ResponseWriter, r *http.Request) {

	country := r.URL.Query().Get("name")

	if country == "" {
		http.Error(w, "Falta el país", http.StatusBadRequest)
		return
	}

	servers, err := database.GetServers()

	if err != nil {
		http.Error(w, "Error obteniendo servidores", http.StatusInternalServerError)
		return
	}

	var result []models.Server

	for _, server := range servers {

		if strings.EqualFold(server.Country, country) {

			result = append(
				result,
				server,
			)

		}

	}

	w.Header().Set("Content-Type", "application/json")

	json.NewEncoder(w).Encode(result)

}

// Buscar por protocolo
func ServersByProtocol(w http.ResponseWriter, r *http.Request) {

	protocol := r.URL.Query().Get("name")

	if protocol == "" {
		http.Error(w, "Falta el protocolo", http.StatusBadRequest)
		return
	}

	servers, err := database.GetServers()

	if err != nil {
		http.Error(w, "Error obteniendo servidores", http.StatusInternalServerError)
		return
	}

	var result []models.Server

	for _, server := range servers {

		if strings.EqualFold(server.Protocol, protocol) {

			result = append(
				result,
				server,
			)

		}

	}

	w.Header().Set("Content-Type", "application/json")

	json.NewEncoder(w).Encode(result)

}

// Buscar por estado
func ServersByStatus(w http.ResponseWriter, r *http.Request) {

	status := r.URL.Query().Get("name")

	if status == "" {
		http.Error(w, "Falta el estado", http.StatusBadRequest)
		return
	}

	servers, err := database.GetServers()

	if err != nil {
		http.Error(w, "Error obteniendo servidores", http.StatusInternalServerError)
		return
	}

	var result []models.Server

	for _, server := range servers {

		if strings.EqualFold(server.Status, status) {

			result = append(
				result,
				server,
			)

		}

	}

	w.Header().Set("Content-Type", "application/json")

	json.NewEncoder(w).Encode(result)

}

// Elegir mejor servidor
func BestServer(w http.ResponseWriter, r *http.Request) {

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
		http.Error(w, "Error obteniendo servidores", http.StatusInternalServerError)
		return
	}

	var best *models.Server

	for i := range servers {

		server := &servers[i]

		if server.Status != "online" {

			continue

		}

		if best == nil {

			best = server

			continue

		}

		if server.Latency < best.Latency {

			best = server

			continue

		}

		if server.Latency == best.Latency &&
			server.CurrentUsers < best.CurrentUsers {

			best = server

		}

	}

	if best == nil {

		http.Error(
			w,
			"No hay servidores disponibles",
			http.StatusNotFound,
		)

		return
	}

	w.Header().Set("Content-Type", "application/json")

	json.NewEncoder(w).Encode(best)

}
