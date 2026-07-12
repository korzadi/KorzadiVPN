package database

import (
	"korzadivpn/models"
)

// GetAdminByEmail busca un administrador por email.
func GetAdminByEmail(email string) (*models.Admin, error) {

	row := DB.QueryRow(`
		SELECT
		id,
		email,
		role,
		created_at
		FROM admins
		WHERE email = ?
		LIMIT 1
	`, email)

	var admin models.Admin

	err := row.Scan(
		&admin.ID,
		&admin.Email,
		&admin.Role,
		&admin.CreatedAt,
	)

	if err != nil {

		return nil, err

	}

	return &admin, nil

}
