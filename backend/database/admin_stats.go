package database


type AdminStats struct {

	TotalUsers int `json:"total_users"`

	ActiveUsers int `json:"active_users"`

	ActiveConnections int `json:"active_connections"`

	OnlineServers int `json:"online_servers"`

	TotalServers int `json:"total_servers"`
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


	return stats, nil
}
