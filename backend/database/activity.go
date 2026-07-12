package database

import (
	"korzadivpn/models"
)

// CreateActivity guarda una actividad.
func CreateActivity(activity models.Activity) error {

	_, err := DB.Exec(`
		INSERT INTO activity_logs
		(email, server, action, device, ip, created_at)
		VALUES (?, ?, ?, ?, ?, ?)
	`,
		activity.Email,
		activity.Server,
		activity.Action,
		activity.Device,
		activity.IP,
		activity.CreatedAt,
	)

	return err
}

// GetActivity obtiene historial de actividad.
func GetActivity() ([]models.Activity, error) {

	activities := make([]models.Activity, 0)

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
		ORDER BY id DESC
	`)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

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
