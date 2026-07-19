package middleware

import (
	"net/http"

	"korzadivpn/internal/database"
)

// AdminAuth protege rutas administrativas.
func AdminAuth(next http.HandlerFunc) http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {

		email, ok := r.Context().
			Value(UserEmailKey).(string)

		if !ok {

			http.Error(
				w,
				"Usuario no autenticado",
				http.StatusUnauthorized,
			)

			return
		}

		admin, err := database.GetAdminByEmail(email)

		if err != nil || admin == nil {

			http.Error(
				w,
				"Acceso administrativo denegado",
				http.StatusForbidden,
			)

			return
		}

		if admin.Role != "superadmin" {

			http.Error(
				w,
				"Permisos insuficientes",
				http.StatusForbidden,
			)

			return
		}

		next(w, r)

	}
}
