package handlers

import (
	"encoding/json"
	"net/http"

	"korzadivpn/database"
)


func AdminDashboard(w http.ResponseWriter, r *http.Request) {


	totalUsers, _ := database.CountUsers()

	activeUsers, _ := database.CountActiveUsers()

	activeConnections, _ := database.CountAllActiveConnections()

	totalServers, _ := database.CountServers()

	onlineServers, _ := database.CountOnlineServers()

	activity, _ := database.GetActivity()



	w.Header().Set(
		"Content-Type",
		"application/json",
	)



	json.NewEncoder(w).Encode(
		map[string]interface{}{

			"users": map[string]interface{}{
				"total": totalUsers,
				"active": activeUsers,
			},

			"servers": map[string]interface{}{
				"total": totalServers,
				"online": onlineServers,
			},

			"connections": map[string]interface{}{
				"active": activeConnections,
			},

			"activity": activity,
		},
	)

}
