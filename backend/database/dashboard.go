package database

// CountUsers cuenta todos los usuarios.
func CountUsers() (int, error) {

	var total int

	err := DB.QueryRow(
		"SELECT COUNT(*) FROM users",
	).Scan(&total)

	return total, err
}

// CountActiveUsers cuenta usuarios activos.
func CountActiveUsers() (int, error) {

	var total int

	err := DB.QueryRow(
		"SELECT COUNT(*) FROM users WHERE status='active'",
	).Scan(&total)

	return total, err
}

// CountAllActiveConnections cuenta conexiones activas.
func CountAllActiveConnections() (int, error) {

	var total int

	err := DB.QueryRow(
		"SELECT COUNT(*) FROM connections WHERE status='connected'",
	).Scan(&total)

	return total, err
}

// CountServers cuenta servidores.
func CountServers() (int, error) {

	var total int

	err := DB.QueryRow(
		"SELECT COUNT(*) FROM servers",
	).Scan(&total)

	return total, err
}

// CountOnlineServers cuenta servidores online.
func CountOnlineServers() (int, error) {

	var total int

	err := DB.QueryRow(
		"SELECT COUNT(*) FROM servers WHERE status='online'",
	).Scan(&total)

	return total, err
}
