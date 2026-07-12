package main

import (
	"fmt"
	"net/http"

	"korzadivpn/database"
	"korzadivpn/routes"
)

func main() {

	database.Connect()

	database.CreateTables()

	database.CreateVPNClientTable()

	database.CreateServers()

	routes.RegisterRoutes()

	fmt.Println("KorzadiVPN API iniciada en puerto 8080 🚀")

	err := http.ListenAndServe(":8080", nil)

	if err != nil {
		fmt.Println("Error:", err)
	}

}
