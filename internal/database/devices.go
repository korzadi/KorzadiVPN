package database

import (
	"database/sql"

	"korzadivpn/internal/models"
)

func RegisterDevice(device models.Device) error {

	_, err := DB.Exec(`
	INSERT INTO devices
	(
		email,
		device_name,
		device_type,
		status,
		last_ip,
		last_server,
		last_seen,
		created_at
	)
	VALUES(?,?,?,?,?,?,?,?)
	`,
		device.Email,
		device.DeviceName,
		device.DeviceType,
		device.Status,
		device.LastIP,
		device.LastServer,
		device.LastSeen,
		device.CreatedAt,
	)

	return err
}

func GetDevicesByEmail(email string) ([]models.Device, error) {

	rows, err := DB.Query(`
	SELECT
	id,
	email,
	device_name,
	device_type,
	status,
	last_ip,
	last_server,
	last_seen,
	created_at
	FROM devices
	WHERE email=?
	ORDER BY id DESC
	`, email)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var devices []models.Device

	for rows.Next() {

		var d models.Device

		err := rows.Scan(
			&d.ID,
			&d.Email,
			&d.DeviceName,
			&d.DeviceType,
			&d.Status,
			&d.LastIP,
			&d.LastServer,
			&d.LastSeen,
			&d.CreatedAt,
		)

		if err != nil {
			return nil, err
		}

		devices = append(devices, d)
	}

	return devices, nil
}

func GetDeviceByName(email, name string) (*models.Device, error) {

	row := DB.QueryRow(`
	SELECT
	id,
	email,
	device_name,
	device_type,
	status,
	last_ip,
	last_server,
	last_seen,
	created_at
	FROM devices
	WHERE email=?
	AND device_name=?
	LIMIT 1
	`,
		email,
		name,
	)

	var d models.Device

	err := row.Scan(
		&d.ID,
		&d.Email,
		&d.DeviceName,
		&d.DeviceType,
		&d.Status,
		&d.LastIP,
		&d.LastServer,
		&d.LastSeen,
		&d.CreatedAt,
	)

	if err == sql.ErrNoRows {
		return nil, nil
	}

	if err != nil {
		return nil, err
	}

	return &d, nil
}

func UpdateDevice(device models.Device) error {

	_, err := DB.Exec(`
	UPDATE devices
	SET
	status=?,
	last_ip=?,
	last_server=?,
	last_seen=?
	WHERE id=?
	AND email=?
	`,
		device.Status,
		device.LastIP,
		device.LastServer,
		device.LastSeen,
		device.ID,
		device.Email,
	)

	return err
}

func UpsertDevice(device models.Device) error {

	old, err := GetDeviceByName(
		device.Email,
		device.DeviceName,
	)

	if err != nil {
		return err
	}

	if old == nil {
		return RegisterDevice(device)
	}

	device.ID = old.ID

	return UpdateDevice(device)
}

func CountActiveDevices(email string) (int, error) {

	row := DB.QueryRow(`
	SELECT COUNT(*)
	FROM devices
	WHERE email=?
	AND status='connected'
	`,
		email)

	var total int

	err := row.Scan(&total)

	return total, err
}

func DeleteDevice(id int, email string) error {

	_, err := DB.Exec(`
	UPDATE connections
	SET status='disconnected'
	WHERE email=?
	AND device=(
		SELECT device_name
		FROM devices
		WHERE id=?
		AND email=?
	)
	`,
		email,
		id,
		email,
	)

	if err != nil {
		return err
	}

	_, err = DB.Exec(`
	DELETE FROM devices
	WHERE id=?
	AND email=?
	`,
		id,
		email,
	)

	return err
}

// GetDeviceByID obtiene un dispositivo específico de un usuario.
func GetDeviceByID(id int, email string) (*models.Device, error) {

	row := DB.QueryRow(`
	SELECT
	id,
	email,
	device_name,
	device_type,
	status,
	last_ip,
	last_server,
	last_seen,
	created_at
	FROM devices
	WHERE id=?
	AND email=?
	LIMIT 1
	`,
		id,
		email,
	)

	var d models.Device

	err := row.Scan(
		&d.ID,
		&d.Email,
		&d.DeviceName,
		&d.DeviceType,
		&d.Status,
		&d.LastIP,
		&d.LastServer,
		&d.LastSeen,
		&d.CreatedAt,
	)

	if err != nil {
		return nil, err
	}

	return &d, nil
}
