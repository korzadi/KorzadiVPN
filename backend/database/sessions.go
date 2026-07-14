package database

import (
	"korzadivpn/models"
)

// CreateSession guarda una nueva sesión activa.
func CreateSession(session models.Session) error {

	_, err := DB.Exec(`
	INSERT INTO sessions
	(
		email,
		token,
		ip,
		device,
		created_at,
		expires_at,
		status
	)
	VALUES(?,?,?,?,?,?,?)
	`,
		session.Email,
		session.Token,
		session.IP,
		session.Device,
		session.CreatedAt,
		session.ExpiresAt,
		session.Status,
	)

	return err
}

// GetActiveSession busca una sesión válida.
func GetActiveSession(token string) (*models.Session, error) {

	row := DB.QueryRow(`
	SELECT
		id,
		email,
		token,
		ip,
		device,
		created_at,
		expires_at,
		status
	FROM sessions
	WHERE token=?
	AND status='active'
	LIMIT 1
	`,
		token,
	)

	var session models.Session

	err := row.Scan(
		&session.ID,
		&session.Email,
		&session.Token,
		&session.IP,
		&session.Device,
		&session.CreatedAt,
		&session.ExpiresAt,
		&session.Status,
	)

	if err != nil {
		return nil, err
	}

	return &session, nil
}

// CloseSession cierra una sesión.
func CloseSession(token string) error {

	_, err := DB.Exec(`
	UPDATE sessions
	SET status='closed'
	WHERE token=?
	`,
		token,
	)

	return err
}

// GetUserSessions obtiene sesiones de un usuario.
func GetUserSessions(email string) ([]models.Session, error) {

	rows, err := DB.Query(`
	SELECT
	id,
	email,
	ip,
	device,
	created_at,
	expires_at,
	status
	FROM sessions
	WHERE email=?
	ORDER BY id DESC
	`,
		email,
	)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var sessions []models.Session

	for rows.Next() {

		var session models.Session

		err := rows.Scan(
			&session.ID,
			&session.Email,
			&session.IP,
			&session.Device,
			&session.CreatedAt,
			&session.ExpiresAt,
			&session.Status,
		)

		if err != nil {
			return nil, err
		}

		sessions = append(
			sessions,
			session,
		)
	}

	return sessions, nil
}

// RevokeSession revoca una sesión activa.
func RevokeSession(token string) error {

	_, err := DB.Exec(`
	UPDATE sessions
	SET status='revoked'
	WHERE token=?
	`,
		token,
	)

	return err
}

// GetActiveSessionsByEmail obtiene sesiones activas del usuario.
func GetActiveSessionsByEmail(email string) ([]models.Session, error) {

	rows, err := DB.Query(`
	SELECT
	id,
	email,
	ip,
	device,
	created_at,
	expires_at,
	status
	FROM sessions
	WHERE email=?
	AND status='active'
	ORDER BY id DESC
	`,
		email,
	)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var sessions []models.Session

	for rows.Next() {

		var session models.Session

		err := rows.Scan(
			&session.ID,
			&session.Email,
			&session.IP,
			&session.Device,
			&session.CreatedAt,
			&session.ExpiresAt,
			&session.Status,
		)

		if err != nil {
			return nil, err
		}

		sessions = append(
			sessions,
			session,
		)
	}

	return sessions, nil
}
