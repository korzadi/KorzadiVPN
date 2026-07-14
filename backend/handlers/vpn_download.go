package handlers

import (
	"net/http"

	"korzadivpn/database"
	"korzadivpn/middleware"
)

func DownloadVPNProfile(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodGet {

		http.Error(
			w,
			"Método no permitido",
			http.StatusMethodNotAllowed,
		)

		return
	}

	email, ok := r.Context().Value(
		middleware.UserEmailKey,
	).(string)

	if !ok {

		http.Error(
			w,
			"Usuario no autenticado",
			http.StatusUnauthorized,
		)

		return
	}

	profile, err := database.GetVPNProfile(email)

	if err != nil {

		http.Error(
			w,
			"Perfil VPN no encontrado",
			http.StatusNotFound,
		)

		return
	}

	filename := "KorzadiVPN-" + profile.Server + ".conf"

	w.Header().Set(
		"Content-Disposition",
		"attachment; filename="+filename,
	)

	w.Header().Set(
		"Content-Type",
		"text/plain",
	)

	w.Header().Set(
		"Cache-Control",
		"no-store",
	)

	w.Write([]byte(profile.Config))

}
