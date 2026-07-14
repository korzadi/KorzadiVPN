package handlers

import (
	"encoding/json"
	"net/http"

	"korzadivpn/database"
	"korzadivpn/middleware"
	"korzadivpn/models"
)

func UserDashboard(
	w http.ResponseWriter,
	r *http.Request,
) {

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

	connection, _ :=
		database.GetActiveConnection(email)

	history, _ :=
		database.GetConnectionsByEmail(email)

	activities, _ :=
		database.GetActivityByEmail(email)

	sessions, _ :=
		database.GetActiveSessionsByEmail(email)

	devicesUsed, _ :=
		database.CountActiveConnections(email)

	deviceLimit :=
		models.GetDeviceLimit(
			user.Plan,
		)

	totalConnections, _ :=
		database.CountUserConnections(email)

	totalDevices, _ :=
		database.CountUserDevices(email)

	lastConnection, _ :=
		database.GetLastConnection(email)

	mostUsedServer, _ :=
		database.GetMostUsedServer(email)

	totalDisconnects, _ :=
		database.CountUserDisconnects(email)

	activeConnections, _ :=
		database.CountUserActiveConnections(email)

	lastDevice, _ :=
		database.GetLastDevice(email)

	connectionRecords, _ :=
		database.GetTotalConnectionTime(email)

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

			"statistics": map[string]interface{}{

				"total_connections": totalConnections,

				"active_connections": activeConnections,

				"total_disconnects": totalDisconnects,

				"total_devices": totalDevices,

				"last_connection": lastConnection,

				"last_device": lastDevice,

				"most_used_server": mostUsedServer,

				"connection_records": connectionRecords,
			},

			"sessions": map[string]interface{}{

				"active": sessions,
			},

			"history": history,

			"activity": activities,
		},
	)

}
