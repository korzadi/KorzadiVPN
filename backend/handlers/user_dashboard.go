package handlers

import (
	"encoding/json"
	"net/http"

	"korzadivpn/database"
	"korzadivpn/middleware"
	"korzadivpn/models"
)

// UserDashboard muestra información del usuario.
func UserDashboard(w http.ResponseWriter, r *http.Request) {

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

	user, err := database.GetUser(email)

	if err != nil {

		http.Error(
			w,
			"Usuario no encontrado",
			http.StatusNotFound,
		)

		return
	}

	connection, _ := database.GetActiveConnection(email)

	activities, _ := database.GetActivityByEmail(email)

	devicesUsed, _ := database.CountActiveConnections(email)

	deviceLimit := models.GetDeviceLimit(
		user.Plan,
	)

	w.Header().Set(
		"Content-Type",
		"application/json",
	)

	json.NewEncoder(w).Encode(
		map[string]interface{}{

			"user": user,

			"plan": map[string]interface{}{

				"name": user.Plan,

				"device_limit": deviceLimit,

				"devices_used": devicesUsed,
			},

			"vpn": connection,

			"activity": activities,
		},
	)

}
