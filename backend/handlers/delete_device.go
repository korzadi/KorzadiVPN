package handlers

import (
	"net/http"
	"strconv"
	"strings"

	"korzadivpn/database"
	"korzadivpn/middleware"
)

func DeleteDevice(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodDelete {
		http.Error(
			w,
			"Metodo no permitido",
			http.StatusMethodNotAllowed,
		)
		return
	}

	email, ok := r.Context().
		Value(middleware.UserEmailKey).(string)

	if !ok {
		http.Error(
			w,
			"Usuario no autenticado",
			http.StatusUnauthorized,
		)
		return
	}

	idText := strings.TrimPrefix(
		r.URL.Path,
		"/api/user/device/",
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

	err = database.DeleteDevice(
		id,
		email,
	)

	if err != nil {
		http.Error(
			w,
			"Error eliminando dispositivo",
			http.StatusInternalServerError,
		)
		return
	}

	w.Header().Set(
		"Content-Type",
		"application/json",
	)

	w.Write([]byte(`{
		"message":"Dispositivo eliminado correctamente"
	}`))
}
