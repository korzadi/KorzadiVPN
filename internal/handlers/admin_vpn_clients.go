package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"
	"strings"

	"korzadivpn/internal/database"
)

func AdminVPNClients(
	w http.ResponseWriter,
	r *http.Request,
) {

	if r.Method != http.MethodGet {

		http.Error(
			w,
			"Metodo no permitido",
			http.StatusMethodNotAllowed,
		)

		return
	}

	clients, err :=
		database.GetAllVPNClients()

	if err != nil {

		http.Error(
			w,
			"Error obteniendo clientes VPN",
			http.StatusInternalServerError,
		)

		return
	}

	w.Header().Set(
		"Content-Type",
		"application/json",
	)

	json.NewEncoder(w).Encode(
		map[string]interface{}{
			"total":   len(clients),
			"clients": clients,
		},
	)

}

func AdminChangeVPNClientStatus(
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

	var req struct {
		ID int `json:"id"`

		Status string `json:"status"`
	}

	err := json.NewDecoder(
		r.Body,
	).Decode(&req)

	if err != nil {

		http.Error(
			w,
			"JSON invalido",
			http.StatusBadRequest,
		)

		return
	}

	err =
		database.UpdateVPNClientStatusByID(
			req.ID,
			req.Status,
		)

	if err != nil {

		http.Error(
			w,
			"Error actualizando cliente",
			http.StatusInternalServerError,
		)

		return
	}

	json.NewEncoder(w).Encode(
		map[string]string{
			"message": "Estado actualizado",
		},
	)

}

func AdminDeleteVPNClient(
	w http.ResponseWriter,
	r *http.Request,
) {

	if r.Method != http.MethodDelete {

		http.Error(
			w,
			"Metodo no permitido",
			http.StatusMethodNotAllowed,
		)

		return
	}

	idText :=
		strings.TrimPrefix(
			r.URL.Path,
			"/api/admin/vpn-client/",
		)

	id, err :=
		strconv.Atoi(idText)

	if err != nil {

		http.Error(
			w,
			"ID invalido",
			http.StatusBadRequest,
		)

		return
	}

	err =
		database.DeleteVPNClientByID(
			id,
		)

	if err != nil {

		http.Error(
			w,
			"Error eliminando cliente",
			http.StatusInternalServerError,
		)

		return
	}

	json.NewEncoder(w).Encode(
		map[string]string{
			"message": "Cliente VPN eliminado",
		},
	)

}
