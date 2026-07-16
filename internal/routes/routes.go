package routes

import (
	"net/http"

	"korzadivpn/internal/handlers"
	"korzadivpn/internal/middleware"
)

func RegisterRoutes(mux *http.ServeMux) {

	mux.HandleFunc("/api/admin/nodes", handlers.ListVPNNodes)

	mux.HandleFunc("/api/admin/nodes/create", handlers.CreateVPNNode)

	mux.HandleFunc("/api/vpn/connect", handlers.StartVPNConnection)

	mux.HandleFunc(
		"/status",
		handlers.Status,
	)

	// USUARIOS

	mux.HandleFunc(
		"/api/register",
		handlers.Register,
	)

	mux.HandleFunc(
		"/api/login",
		handlers.Login,
	)

	mux.HandleFunc(
		"/api/profile",
		middleware.Auth(
			handlers.Profile,
		),
	)

	mux.HandleFunc(
		"/api/user/dashboard",
		middleware.Auth(
			handlers.UserDashboard,
		),
	)

	mux.HandleFunc(
		"/api/user/devices",
		middleware.Auth(
			handlers.UserDevices,
		),
	)

	mux.HandleFunc(
		"/api/user/device/",
		middleware.Auth(
			handlers.DeleteDevice,
		),
	)

	// SERVIDORES

	mux.HandleFunc(
		"/api/servers",
		handlers.Servers,
	)

	mux.HandleFunc(
		"/api/servers/best",
		handlers.BestServer,
	)

	mux.HandleFunc(
		"/api/servers/load",
		handlers.ServersLoad,
	)

	mux.HandleFunc(
		"/api/plans",
		handlers.Plans,
	)

	mux.HandleFunc(
		"/api/change-plan",
		middleware.Auth(
			handlers.ChangePlan,
		),
	)

	// VPN

	mux.HandleFunc(
		"/api/connect",
		middleware.Auth(
			handlers.Connect,
		),
	)

	mux.HandleFunc(
		"/api/disconnect",
		middleware.Auth(
			handlers.Disconnect,
		),
	)

	mux.HandleFunc(
		"/api/vpn/profile",
		middleware.Auth(
			handlers.CreateVPNProfile,
		),
	)

	mux.HandleFunc(
		"/api/vpn/profile/get",
		middleware.Auth(
			handlers.GetVPNProfile,
		),
	)

	mux.HandleFunc(
		"/api/vpn/profile/download",
		middleware.Auth(
			handlers.DownloadVPNProfile,
		),
	)

	// NUEVO ESTADO VPN

	mux.HandleFunc(
		"/api/vpn/status",
		middleware.Auth(
			handlers.VPNStatus,
		),
	)

	// CLIENTE VPN WIREGUARD

	mux.HandleFunc(
		"/api/vpn/client/create",
		middleware.Auth(
			handlers.CreateVPNClient,
		),
	)

	mux.HandleFunc(
		"/api/vpn/config",
		middleware.Auth(
			handlers.CreateVPNConfig,
		),
	)

	// ADMIN

	mux.HandleFunc(
		"/api/admin/users",
		middleware.Auth(
			middleware.AdminAuth(
				handlers.AdminUsers,
			),
		),
	)

	mux.HandleFunc(
		"/api/admin/servers",
		middleware.Auth(
			middleware.AdminAuth(
				handlers.AdminServers,
			),
		),
	)

	mux.HandleFunc(
		"/api/admin/connections",
		middleware.Auth(
			middleware.AdminAuth(
				handlers.AdminConnections,
			),
		),
	)

	mux.HandleFunc(
		"/api/admin/stats",
		middleware.Auth(
			middleware.AdminAuth(
				handlers.AdminStats,
			),
		),
	)

	mux.HandleFunc(
		"/api/admin/activity",
		middleware.Auth(
			middleware.AdminAuth(
				handlers.AdminActivity,
			),
		),
	)

	mux.HandleFunc(
		"/api/admin/dashboard",
		middleware.Auth(
			middleware.AdminAuth(
				handlers.AdminDashboard,
			),
		),
	)

}
