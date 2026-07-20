package handlers

import (
	"encoding/json"
	"log"
	"net/http"
	"time"

	"korzadivpn/internal/database"
	"korzadivpn/internal/middleware"
	"korzadivpn/internal/models"
	"korzadivpn/internal/services"
	"korzadivpn/pkg/utils"
	"korzadivpn/vpn-core/generator"
	"korzadivpn/vpn-core/wireguard"
)

func CreateVPNClient(
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

	email, ok :=
		r.Context().
			Value(middleware.UserEmailKey).(string)

	if !ok {

		http.Error(
			w,
			"Usuario no autenticado",
			http.StatusUnauthorized,
		)

		return
	}

	server, err :=
		database.GetBestServer()

	if err != nil {
		log.Printf("Error al obtener el mejor servidor: %v", err)
		http.Error(
			w,
			"No hay servidores disponibles",
			http.StatusInternalServerError,
		)

		return
	}

	clientIP, err :=
		database.GetNextVPNClientIP()

	if err != nil {
		log.Printf("Error al obtener IP: %v", err)
		http.Error(
			w,
			"No hay IP disponible",
			http.StatusInternalServerError,
		)

		return
	}

	publicKey, privateKey :=
		utils.GenerateWireGuardKeys()

	now :=
		time.Now().
			UTC().
			Format(time.RFC3339)

		// Generar configuración WireGuard del cliente
	wgPeer := wireguard.Peer{
		PrivateKey: privateKey,
		IPAddress:  clientIP,
	}

	wgServer := wireguard.Server{
		PublicKey: server.ServerPublicKey,
		Endpoint:  server.ServerIP,
		Port:      server.WireGuardPort,
	}

	clientConfig := generator.GenerateClientConfig(
		wgPeer,
		wgServer,
	)

	client := models.VPNClient{

		Email: email,

		ServerID: server.ID,

		NodeID: server.ID,

		ClientName: "Korzadi-Device",

		DeviceName: "Korzadi Device",

		DeviceType: "WireGuard",

		ClientIP: clientIP,

		PublicKey: publicKey,

		PrivateKey: privateKey,

		Protocol: "wireguard",

		DNS: "1.1.1.1",

		MTU: 1420,

		AllowedIPs: "0.0.0.0/0, ::/0",
		Config:     clientConfig,

		Endpoint: server.ServerIP,

		Status: "pending",

		ConnectionStatus: "offline",

		Plan: "free",

		BandwidthLimit: 0,

		DataUsed: 0,

		MaxDevices: 1,

		CreatedAt: now,

		UpdatedAt: now,
	}

	// Reintentar creación de cliente hasta 3 veces si hay conflicto de IP
	success := false
	for i := 0; i < 3; i++ {
		err = database.CreateVPNClient(client)
		if err == nil {
			success = true
			break
		}
		// Registrar error y reintentar si es posible
		log.Printf("Intento %d fallido al crear cliente PENDING para IP %s: %v", i+1, client.ClientIP, err)

		// Obtener nueva IP para el próximo intento si falla
		newIP, newErr := database.GetNextVPNClientIP()
		if newErr != nil {
			break
		}
		client.ClientIP = newIP
	}

	if !success {
		log.Printf("Error final creando cliente tras 3 intentos: %v", err)
		http.Error(
			w,
			"Error creando cliente VPN tras reintentos",
			http.StatusInternalServerError,
		)

		return
	}

	// Provisionar en WireGuard
	vpnService := services.NewVPNService()
	err = vpnService.ProvisionPeer(&client)

	if err != nil {
		log.Printf("Error provisionando peer para %s: %v", client.Email, err)
		database.UpdateVPNClientStatusByEmail(client.Email, "failed")

		// Rollback opcional
		vpnService.WireGuard.RemovePeer(client.PublicKey)

		http.Error(w, "Error provisionando VPN", http.StatusInternalServerError)
		return
	}

	database.UpdateVPNClientStatusByEmail(client.Email, "active")
	client.Status = "active"

	client.Config = clientConfig

	database.IncrementServerUsers(
		server.ID,
	)

	device := models.Device{

		Email: email,

		DeviceName: client.DeviceName,

		DeviceType: client.DeviceType,

		Status: "active",

		LastServer: server.Name,

		LastSeen: now,

		CreatedAt: now,
	}

	err =
		database.UpsertDevice(
			device,
		)

	if err != nil {

		http.Error(
			w,
			"Error registrando dispositivo",
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

			"message": "Cliente VPN creado correctamente",

			"client": client,

			"server": server,

			"device": device,
		},
	)

}
