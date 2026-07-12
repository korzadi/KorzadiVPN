package database

import (
	"database/sql"

	"korzadivpn/models"
)

// RegisterDevice registra un dispositivo.
func RegisterDevice(device models.Device) error {

	_, err := DB.Exec(`
	INSERT INTO devices
	(email,device_name,device_type,status,last_ip,last_server,last_seen,created_at)
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

// GetDevicesByEmail obtiene dispositivos del usuario.
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

// GetDeviceByName busca un dispositivo.
func GetDeviceByName(email, deviceName string) (*models.Device, error) {

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
	`, email, deviceName)

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

// DeleteDevice elimina un dispositivo del usuario.
func DeleteDevice(id int, email string) error {

	_, err := DB.Exec(`
	DELETE FROM devices
	WHERE id=?
	AND email=?
	`,
		id,
		email,
	)

	return err
}

// UpdateDevice actualiza la información de un dispositivo.
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
