package main

import (
	"fmt"
	"net/http"

	"korzadivpn/internal/database"
	"korzadivpn/internal/middleware"
	"korzadivpn/internal/routes"
)

func corsMiddleware(next http.Handler) http.Handler {

	return http.HandlerFunc(func(
		w http.ResponseWriter,
		r *http.Request,
	) {

		w.Header().Set(
			"Access-Control-Allow-Origin",
			"*",
		)

		w.Header().Set(
			"Access-Control-Allow-Methods",
			"GET, POST, DELETE, PUT, OPTIONS",
		)

		w.Header().Set(
			"Access-Control-Allow-Headers",
			"Content-Type, Authorization",
		)

		if r.Method == http.MethodOptions {

			w.WriteHeader(http.StatusOK)

			return
		}

		next.ServeHTTP(w, r)

	})

}

func main() {

	database.Connect()

	database.ConfigureSQLite()

	database.CreateTables()

	database.CreateVPNClientTable()

	database.MigrateVPNClients()

	database.MigrateServerKeys()

	err := database.CreateServers()

	if err != nil {

		fmt.Println(
			"Error creando servidores:",
			err,
		)
	}

	mux := http.NewServeMux()
	routes.RegisterRoutes(mux)

	// Envolver en middleware
	handler := middleware.SecurityHeaders(
		middleware.RateLimiter(
			corsMiddleware(mux),
		),
	)

	fmt.Println(
		"KorzadiVPN API iniciada en puerto 8080 🚀",
	)

	err = http.ListenAndServe(
		":8080",
		handler,
	)

	if err != nil {

		fmt.Println(
			"Error:",
			err,
		)
	}

}
