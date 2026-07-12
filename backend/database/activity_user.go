package database

import (
	"korzadivpn/models"
)

// GetActivityByEmail obtiene actividad de un usuario.
func GetActivityByEmail(email string) ([]models.Activity, error) {

	rows, err := DB.Query(`
		SELECT
			id,
			email,
			server,
			action,
			device,
			ip,
			created_at
		FROM activity_logs
		WHERE email = ?
		ORDER BY id DESC
	`, email)

	if err != nil {

		return nil, err

	}

	defer rows.Close()

	var activities []models.Activity

	for rows.Next() {

		var activity models.Activity

		err := rows.Scan(
			&activity.ID,
			&activity.Email,
			&activity.Server,
			&activity.Action,
			&activity.Device,
			&activity.IP,
			&activity.CreatedAt,
		)

		if err != nil {

			return nil, err

		}

		activities = append(
			activities,
			activity,
		)

	}

	return activities, nil

}
