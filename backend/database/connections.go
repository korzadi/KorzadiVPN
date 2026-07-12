package database

import (
	"database/sql"

	"korzadivpn/models"
)

// CreateConnection guarda una nueva conexión VPN.
func CreateConnection(connection models.Connection) error {

	_, err := DB.Exec(`
	INSERT INTO connections
	(email,server_id,server,status,device,ip,connected_at)
	VALUES(?,?,?,?,?,?,?)
	`,
		connection.Email,
		connection.ServerID,
		connection.Server,
		connection.Status,
		connection.Device,
		connection.IP,
		connection.ConnectedAt,
	)

	return err
}

// GetConnectionsByEmail obtiene conexiones del usuario.
func GetConnectionsByEmail(email string) ([]models.Connection, error) {

	rows, err := DB.Query(`
	SELECT
	email,
	server_id,
	server,
	status,
	device,
	ip,
	connected_at
	FROM connections
	WHERE email=?
	`, email)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var connections []models.Connection

	for rows.Next() {

		var c models.Connection

		err := rows.Scan(
			&c.Email,
			&c.ServerID,
			&c.Server,
			&c.Status,
			&c.Device,
			&c.IP,
			&c.ConnectedAt,
		)

		if err != nil {
			return nil, err
		}

		connections = append(connections, c)
	}

	return connections, nil
}

// GetActiveConnection obtiene una conexión activa.
func GetActiveConnection(email string) (*models.Connection, error) {

	row := DB.QueryRow(`
	SELECT
	email,
	server_id,
	server,
	status,
	device,
	ip,
	connected_at
	FROM connections
	WHERE email=?
	AND status='connected'
	ORDER BY id DESC
	LIMIT 1
	`, email)

	var c models.Connection

	err := row.Scan(
		&c.Email,
		&c.ServerID,
		&c.Server,
		&c.Status,
		&c.Device,
		&c.IP,
		&c.ConnectedAt,
	)

	if err == sql.ErrNoRows {
		return nil, nil
	}

	if err != nil {
		return nil, err
	}

	return &c, nil
}

// CountActiveConnections cuenta DISPOSITIVOS únicos activos.
func CountActiveConnections(email string) (int, error) {

	row := DB.QueryRow(`
	SELECT COUNT(DISTINCT device)
	FROM connections
	WHERE email=?
	AND status='connected'
	AND device IS NOT NULL
	AND device <> ''
	`, email)

	var total int

	err := row.Scan(&total)

	return total, err
}

// DisconnectConnection desconecta usuario.
func DisconnectConnection(email string) error {

	_, err := DB.Exec(`
	UPDATE connections
	SET status='disconnected'
	WHERE email=?
	AND status='connected'
	`, email)

	return err
}

// GetAllActiveConnections obtiene conexiones activas para admin.
func GetAllActiveConnections() ([]models.Connection, error) {

	rows, err := DB.Query(`
	SELECT
	email,
	server_id,
	server,
	status,
	device,
	ip,
	connected_at
	FROM connections
	WHERE status='connected'
	ORDER BY id DESC
	`)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var connections []models.Connection

	for rows.Next() {

		var c models.Connection

		err := rows.Scan(
			&c.Email,
			&c.ServerID,
			&c.Server,
			&c.Status,
			&c.Device,
			&c.IP,
			&c.ConnectedAt,
		)

		if err != nil {
			return nil, err
		}

		connections = append(connections, c)
	}

	return connections, nil
}
