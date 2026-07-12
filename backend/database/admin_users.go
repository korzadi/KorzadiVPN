package database

import (
	"korzadivpn/models"
)

// GetUsers obtiene usuarios sin mostrar contraseñas.
func GetUsers() ([]models.User, error) {

	rows, err := DB.Query(`
		SELECT
		email,
		plan,
		status
		FROM users
	`)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var users []models.User

	for rows.Next() {

		var user models.User

		err := rows.Scan(
			&user.Email,
			&user.Plan,
			&user.Status,
		)

		if err != nil {
			return nil, err
		}

		users = append(
			users,
			user,
		)
	}

	return users, nil
}
