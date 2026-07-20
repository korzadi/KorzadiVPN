package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"korzadivpn/internal/database"
	"korzadivpn/internal/middleware"
	"korzadivpn/internal/models"
)

func CreateVPNProfile(
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

	client, err :=
		database.GetVPNClientByEmail(email)

	if err != nil {

		http.Error(
			w,
			"No existe cliente VPN",
			http.StatusNotFound,
		)

		return
	}

	server, err :=
		database.GetServerByID(
			client.ServerID,
		)

	if err != nil {

		http.Error(
			w,
			"Servidor no encontrado",
			http.StatusNotFound,
		)

		return
	}

	config :=
		"[Interface]\n" +
			"PrivateKey = " + client.PrivateKey + "\n" +
			"Address = " + client.ClientIP + "/24\n" +
			"DNS = " + server.DNS + "\n\n" +
			"[Peer]\n" +
			"PublicKey = " + server.ServerPublicKey + "\n" +
			"Endpoint = " +
			server.ServerIP + ":" +
			strconv.Itoa(server.WireGuardPort) +
			"\n" +
			"AllowedIPs = 0.0.0.0/0\n" +
			"PersistentKeepalive = 25"

	profile :=
		models.VPNProfile{

			Email: email,

			ServerID: server.ID,

			Server: server.Name,

			Protocol: server.Protocol,

			PublicKey: client.PublicKey,

			PrivateKey: client.PrivateKey,

			Config: config,

			Status: "ready",
		}

	err =
		database.CreateVPNProfile(
			profile,
		)

	if err != nil {

		http.Error(
			w,
			"Error creando perfil",
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

			"message": "Perfil VPN creado correctamente",

			"server": server.Name,

			"protocol": server.Protocol,

			"status": "profile.Status",

			"public_key": client.PublicKey,
		},
	)

}

func GetVPNProfile(
	w http.ResponseWriter,
	r *http.Request,
) {

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

	profile, err :=
		database.GetVPNProfile(email)

	if err != nil {

		http.Error(
			w,
			"Perfil VPN no encontrado",
			http.StatusNotFound,
		)

		return
	}

	w.Header().Set(
		"Content-Type",
		"application/json",
	)

	json.NewEncoder(w).Encode(
		map[string]interface{}{

			"email": profile.Email,

			"server": profile.Server,

			"protocol": profile.Protocol,

			"public_key": profile.PublicKey,

			"status": profile.Status,
		},
	)

}
