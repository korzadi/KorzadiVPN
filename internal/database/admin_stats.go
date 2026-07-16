package database

type AdminStats struct {
	TotalUsers int `json:"total_users"`

	ActiveUsers int `json:"active_users"`

	ActiveConnections int `json:"active_connections"`

	TotalConnections int `json:"total_connections"`

	TotalDevices int `json:"total_devices"`

	OnlineServers int `json:"online_servers"`

	TotalServers int `json:"total_servers"`

	LastActivity string `json:"last_activity"`
}

// GetAdminStats obtiene estadísticas generales.
func GetAdminStats() (AdminStats, error) {

	var stats AdminStats

	err := DB.QueryRow(`
	SELECT COUNT(*)
	FROM users
	`).Scan(
		&stats.TotalUsers,
	)

	if err != nil {
		return stats, err
	}

	err = DB.QueryRow(`
	SELECT COUNT(*)
	FROM users
	WHERE status='active'
	`).Scan(
		&stats.ActiveUsers,
	)

	if err != nil {
		return stats, err
	}

	err = DB.QueryRow(`
	SELECT COUNT(*)
	FROM connections
	WHERE status='connected'
	`).Scan(
		&stats.ActiveConnections,
	)

	if err != nil {
		return stats, err
	}

	err = DB.QueryRow(`
	SELECT COUNT(*)
	FROM connections
	`).Scan(
		&stats.TotalConnections,
	)

	if err != nil {
		return stats, err
	}

	err = DB.QueryRow(`
	SELECT COUNT(*)
	FROM devices
	`).Scan(
		&stats.TotalDevices,
	)

	if err != nil {
		return stats, err
	}

	err = DB.QueryRow(`
	SELECT COUNT(*)
	FROM servers
	WHERE status='online'
	`).Scan(
		&stats.OnlineServers,
	)

	if err != nil {
		return stats, err
	}

	err = DB.QueryRow(`
	SELECT COUNT(*)
	FROM servers
	`).Scan(
		&stats.TotalServers,
	)

	if err != nil {
		return stats, err
	}

	DB.QueryRow(`
	SELECT created_at
	FROM activity_logs
	ORDER BY id DESC
	LIMIT 1
	`).Scan(
		&stats.LastActivity,
	)

	return stats, nil
}
