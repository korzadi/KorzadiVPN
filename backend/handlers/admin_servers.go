package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"
	"strings"

	"korzadivpn/database"
	"korzadivpn/models"
)

func AdminServers(
	w http.ResponseWriter,
	r *http.Request,
) {

	switch r.Method {

	case http.MethodGet:

		// Detalle por ID
		if strings.HasPrefix(
			r.URL.Path,
			"/api/admin/servers/",
		) {

			idText := strings.TrimPrefix(
				r.URL.Path,
				"/api/admin/servers/",
			)

			id, err := strconv.Atoi(idText)

			if err != nil {

				http.Error(
					w,
					"ID invalido",
					http.StatusBadRequest,
				)

				return
			}

			server, err := database.GetServerAdminByID(id)

			if err != nil {

				http.Error(
					w,
					"Servidor no encontrado",
					http.StatusNotFound,
				)

				return
			}

			w.Header().Set(
				"Content-Type",
				"application/json",
			)

			json.NewEncoder(w).Encode(server)

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

		w.Header().Set(
			"Content-Type",
			"application/json",
		)

		json.NewEncoder(w).Encode(
			map[string]interface{}{
				"total":   len(servers),
				"servers": servers,
			},
		)

	case http.MethodPost:

		var server models.Server

		err := json.NewDecoder(
			r.Body,
		).Decode(&server)

		if err != nil {

			http.Error(
				w,
				"JSON invalido",
				http.StatusBadRequest,
			)

			return
		}

		if server.Name == "" ||
			server.Country == "" ||
			server.ServerIP == "" {

			http.Error(
				w,
				"Datos incompletos",
				http.StatusBadRequest,
			)

			return
		}

		err = database.CreateServer(server)

		if err != nil {

			http.Error(
				w,
				"Error creando servidor",
				http.StatusInternalServerError,
			)

			return
		}

		json.NewEncoder(w).Encode(
			map[string]string{
				"message": "Servidor creado correctamente",
			},
		)

	case http.MethodPut:

		var server models.Server

		err := json.NewDecoder(
			r.Body,
		).Decode(&server)

		if err != nil {

			http.Error(
				w,
				"JSON invalido",
				http.StatusBadRequest,
			)

			return
		}

		if server.ID == 0 {

			http.Error(
				w,
				"ID requerido",
				http.StatusBadRequest,
			)

			return
		}

		err = database.UpdateServer(server)

		if err != nil {

			http.Error(
				w,
				"Error actualizando servidor",
				http.StatusInternalServerError,
			)

			return
		}

		json.NewEncoder(w).Encode(
			map[string]string{
				"message": "Servidor actualizado correctamente",
			},
		)

	case http.MethodDelete:

		idText := strings.TrimPrefix(
			r.URL.Path,
			"/api/admin/servers/",
		)

		id, err := strconv.Atoi(idText)

		if err != nil {

			http.Error(
				w,
				"ID invalido",
				http.StatusBadRequest,
			)

			return
		}

		err = database.DeleteServer(id)

		if err != nil {

			http.Error(
				w,
				err.Error(),
				http.StatusBadRequest,
			)

			return
		}

		json.NewEncoder(w).Encode(
			map[string]string{
				"message": "Servidor eliminado correctamente",
			},
		)

	default:

		http.Error(
			w,
			"Metodo no permitido",
			http.StatusMethodNotAllowed,
		)

	}

}
