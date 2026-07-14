package database

import (
	"database/sql"
	"time"

	"korzadivpn/models"
)

// CreateConnection guarda una nueva conexión VPN.
func CreateConnection(connection models.Connection) error {

	_, err := DB.Exec(`
        INSERT INTO connections
        (email,server_id,server,status,device,client_id,ip,connected_at,last_ping)
        VALUES(?,?,?,?,?,?,?,?,?)
        `,
		connection.Email,
		connection.ServerID,
		connection.Server,
		connection.Status,
		connection.Device,
		connection.ClientID,
		connection.IP,
		connection.ConnectedAt,
		connection.LastPing,
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
	client_id,
	ip,
	connected_at,
	disconnected_at,
	last_ping
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
			&c.ClientID,
			&c.IP,
			&c.ConnectedAt,
			&c.DisconnectedAt,
			&c.LastPing,
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
        connected_at,
        client_id,
        disconnected_at,
        last_ping
        FROM connections
        WHERE email=?
        AND status='connected'
        ORDER BY id DESC
        LIMIT 1
        `, email)

	var c models.Connection
	var disconnectedAt sql.NullString
	var lastPing sql.NullString

	err := row.Scan(
		&c.Email,
		&c.ServerID,
		&c.Server,
		&c.Status,
		&c.Device,
		&c.IP,
		&c.ConnectedAt,
		&c.ClientID,
		&disconnectedAt,
		&lastPing,
	)

	if err == sql.ErrNoRows {
		return nil, nil
	}

	if err != nil {
		return nil, err
	}

	if disconnectedAt.Valid {
		c.DisconnectedAt = disconnectedAt.String
	}

	if lastPing.Valid {
		c.LastPing = lastPing.String
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

	now := time.Now().UTC().Format(time.RFC3339)

	_, err := DB.Exec(`
        UPDATE connections
        SET 
        status='disconnected',
        disconnected_at=?,
        last_ping=?
        WHERE email=?
        AND status='connected'
        `,
		now,
		now,
		email,
	)

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
        client_id,
        ip,
        connected_at,
        disconnected_at,
        last_ping
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
			&c.ClientID,
			&c.IP,
			&c.ConnectedAt,
			&c.LastPing,
		)

		if err != nil {
			return nil, err
		}

		connections = append(
			connections,
			c,
		)
	}

	return connections, nil
}

func DisconnectDevice(
	email string,
	device string,
) error {

	now := time.Now().UTC().Format(time.RFC3339)

	_, err := DB.Exec(`
        UPDATE connections
        SET
        status='disconnected',
        disconnected_at=?,
        last_ping=?
        WHERE email=?
        AND device=?
        AND status='connected'
        `,
		now,
		now,
		email,
		device,
	)

	return err
}

// DisconnectDeviceConnection desconecta un dispositivo específico.
func DisconnectDeviceConnection(email string, device string) error {

	now := time.Now().UTC().Format(time.RFC3339)

	_, err := DB.Exec(`
        UPDATE connections
        SET 
        status='disconnected',
        disconnected_at=?,
        last_ping=?
        WHERE email=?
        AND device=?
        AND status='connected'
        `,
		now,
		now,
		email,
		device,
	)

	return err
}

// UpdateLastPing actualiza el último heartbeat del cliente.
func UpdateLastPing(clientID string) error {

	now := time.Now().UTC().Format(time.RFC3339)

	_, err := DB.Exec(`
	UPDATE connections
	SET last_ping=?
	WHERE client_id=?
	AND status='connected'
	`,
		now,
		clientID,
	)

	return err
}

// GetConnectionByClientID valida una conexión por cliente y usuario.
func GetConnectionByClientID(clientID string, email string) (*models.Connection, error) {

	row := DB.QueryRow(`
	SELECT
	email,
	server_id,
	server,
	status,
	device,
	client_id,
	ip,
	connected_at,
	disconnected_at,
	last_ping
	FROM connections
	WHERE client_id=?
	AND email=?
	AND status='connected'
	LIMIT 1
	`,
		clientID,
		email,
	)

	var c models.Connection
	var disconnectedAt sql.NullString
	var lastPing sql.NullString

	err := row.Scan(
		&c.Email,
		&c.ServerID,
		&c.Server,
		&c.Status,
		&c.Device,
		&c.ClientID,
		&c.IP,
		&c.ConnectedAt,
		&disconnectedAt,
		&lastPing,
	)

	if err == sql.ErrNoRows {
		return nil, nil
	}

	if err != nil {
		return nil, err
	}

	if disconnectedAt.Valid {
		c.DisconnectedAt = disconnectedAt.String
	}

	if lastPing.Valid {
		c.LastPing = lastPing.String
	}

	return &c, nil
}
