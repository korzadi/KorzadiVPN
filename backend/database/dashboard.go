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

// CountUserConnections cuenta conexiones históricas del usuario.
func CountUserConnections(email string) (int, error) {

	var total int

	err := DB.QueryRow(
		"SELECT COUNT(*) FROM connections WHERE email=?",
		email,
	).Scan(&total)

	return total, err
}

// CountUserDevices cuenta dispositivos del usuario.
func CountUserDevices(email string) (int, error) {

	var total int

	err := DB.QueryRow(
		"SELECT COUNT(*) FROM devices WHERE email=?",
		email,
	).Scan(&total)

	return total, err
}

// GetLastConnection obtiene última conexión.
func GetLastConnection(email string) (string, error) {

	var date string

	err := DB.QueryRow(`
	SELECT connected_at
	FROM connections
	WHERE email=?
	ORDER BY id DESC
	LIMIT 1
	`,
		email).Scan(&date)

	return date, err
}

// GetMostUsedServer servidor más usado.
func GetMostUsedServer(email string) (string, error) {

	var server string

	err := DB.QueryRow(`
	SELECT server
	FROM connections
	WHERE email=?
	GROUP BY server
	ORDER BY COUNT(*) DESC
	LIMIT 1
	`,
		email).Scan(&server)

	return server, err
}

// CountUserDisconnects cuenta desconexiones del usuario.
func CountUserDisconnects(email string) (int, error) {

	var total int

	err := DB.QueryRow(`
	SELECT COUNT(*)
	FROM connections
	WHERE email=?
	AND status='disconnected'
	`,
		email).Scan(&total)

	return total, err
}

// CountUserActiveConnections cuenta conexiones actuales.
func CountUserActiveConnections(email string) (int, error) {

	var total int

	err := DB.QueryRow(`
	SELECT COUNT(*)
	FROM connections
	WHERE email=?
	AND status='connected'
	`,
		email).Scan(&total)

	return total, err
}

// GetLastDevice obtiene el último dispositivo usado.
func GetLastDevice(email string) (string, error) {

	var device string

	err := DB.QueryRow(`
	SELECT device
	FROM connections
	WHERE email=?
	ORDER BY id DESC
	LIMIT 1
	`,
		email).Scan(&device)

	return device, err
}

// GetTotalConnectionTime obtiene conexiones con tiempo registrado.
func GetTotalConnectionTime(email string) (int, error) {

	var total int

	err := DB.QueryRow(`
	SELECT COUNT(*)
	FROM connections
	WHERE email=?
	AND connected_at IS NOT NULL
	`,
		email).Scan(&total)

	return total, err
}
