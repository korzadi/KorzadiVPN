package database

type AdminUser struct {
	Email string `json:"email"`

	Plan string `json:"plan"`

	Status string `json:"status"`

	Devices int `json:"devices"`

	Connections int `json:"connections"`
}

// GetUsers obtiene usuarios para administración.
func GetUsers() ([]AdminUser, error) {

	rows, err := DB.Query(`
	SELECT
		email,
		plan,
		status
	FROM users
	ORDER BY email
	`)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var users []AdminUser

	for rows.Next() {

		var user AdminUser

		err := rows.Scan(
			&user.Email,
			&user.Plan,
			&user.Status,
		)

		if err != nil {
			return nil, err
		}

		DB.QueryRow(`
		SELECT COUNT(*)
		FROM devices
		WHERE email=?
		`,
			user.Email,
		).Scan(
			&user.Devices,
		)

		DB.QueryRow(`
		SELECT COUNT(*)
		FROM connections
		WHERE email=?
		`,
			user.Email,
		).Scan(
			&user.Connections,
		)

		users = append(
			users,
			user,
		)
	}

	return users, nil
}
