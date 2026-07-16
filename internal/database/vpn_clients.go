package database

import (
	"korzadivpn/internal/models"
	"korzadivpn/pkg/utils"
	"log"
)

func CreateVPNClient(client models.VPNClient) error {
	encryptedKey, err := utils.Encrypt(client.PrivateKey)
	if err != nil {
		return err
	}
	client.PrivateKey = encryptedKey

	_, err = DB.Exec(
		`  
	INSERT INTO vpn_clients  
	(  
		email,  
		server_id,  
		node_id,  
		client_name,  
		device_id,  
		device_name,  
		device_type,  
		client_ip,  
		ipv6,  
		public_key,  
		private_key,  
		preshared_key,  
		config,  
		protocol,  
		dns,  
		mtu,  
		allowed_ips,  
		endpoint,  
		status,  
		connection_status,  
		plan,  
		bandwidth_limit,  
		data_used,  
		max_devices,  
		last_handshake,  
		last_connected,  
		last_disconnected,  
		last_ip,  
		country,  
		city,  
		expires_at,  
		revoked_at,  
		created_at,  
		updated_at  
	)  
	VALUES (?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?)  
	`,
		client.Email,
		client.ServerID,
		client.NodeID,
		client.ClientName,
		client.DeviceID,
		client.DeviceName,
		client.DeviceType,
		client.ClientIP,
		client.IPv6,
		client.PublicKey,
		client.PrivateKey,
		client.PresharedKey,
		client.Config,
		client.Protocol,
		client.DNS,
		client.MTU,
		client.AllowedIPs,
		client.Endpoint,
		client.Status,
		client.ConnectionStatus,
		client.Plan,
		client.BandwidthLimit,
		client.DataUsed,
		client.MaxDevices,
		client.LastHandshake,
		client.LastConnected,
		client.LastDisconnected,
		client.LastIP,
		client.Country,
		client.City,
		client.ExpiresAt,
		client.RevokedAt,
		client.CreatedAt,
		client.UpdatedAt,
	)

	return err

}

func GetVPNClientByEmail(email string) (*models.VPNClient, error) {

	row := DB.QueryRow(
		`  
	SELECT  
		id,  
		email,  
		server_id,  
		node_id,  
		client_name,  
		device_id,  
		device_name,  
		device_type,  
		client_ip,  
		ipv6,  
		public_key,  
		private_key,  
		preshared_key,  
		config,  
		protocol,  
		dns,  
		mtu,  
		allowed_ips,  
		endpoint,  
		status,  
		connection_status,  
		plan,  
		bandwidth_limit,  
		data_used,  
		max_devices,  
		last_handshake,  
		last_connected,  
		last_disconnected,  
		last_ip,  
		country,  
		city,  
		expires_at,  
		revoked_at,  
		created_at,  
		updated_at  
	FROM vpn_clients  
	WHERE email=?  
	AND status='active'  
	ORDER BY id DESC  
	LIMIT 1  
	`,
		email,
	)

	var client models.VPNClient

	err := row.Scan(
		&client.ID,
		&client.Email,
		&client.ServerID,
		&client.NodeID,
		&client.ClientName,
		&client.DeviceID,
		&client.DeviceName,
		&client.DeviceType,
		&client.ClientIP,
		&client.IPv6,
		&client.PublicKey,
		&client.PrivateKey,
		&client.PresharedKey,
		&client.Config,
		&client.Protocol,
		&client.DNS,
		&client.MTU,
		&client.AllowedIPs,
		&client.Endpoint,
		&client.Status,
		&client.ConnectionStatus,
		&client.Plan,
		&client.BandwidthLimit,
		&client.DataUsed,
		&client.MaxDevices,
		&client.LastHandshake,
		&client.LastConnected,
		&client.LastDisconnected,
		&client.LastIP,
		&client.Country,
		&client.City,
		&client.ExpiresAt,
		&client.RevokedAt,
		&client.CreatedAt,
		&client.UpdatedAt,
	)

	if err != nil {
		return nil, err
	}

	// Intenta descifrar, si falla, asume texto plano (compatibilidad con datos antiguos)
	decryptedKey, err := utils.Decrypt(client.PrivateKey)
	if err == nil {
		client.PrivateKey = decryptedKey
	} else {
		log.Printf("Advertencia: No se pudo descifrar la clave privada para %s, asumiendo texto plano: %v", client.Email, err)
	}

	return &client, nil

}

func DeleteVPNClient(id int, email string) error {

	_, err := DB.Exec(
		`  
	DELETE FROM vpn_clients  
	WHERE id=? AND email=?  
	`,
		id,
		email,
	)

	return err

}

func UpdateVPNClientConfig(
	id int,
	config string,
) error {

	_, err := DB.Exec(
		`  
	UPDATE vpn_clients  
	SET  
		config=?,  
		updated_at=datetime('now')  
	WHERE id=?  
	`,
		config,
		id,
	)

	return err

}

func UpdateVPNClientStatusByEmail(email string, status string) error {
	_, err := DB.Exec(
		`
		UPDATE vpn_clients
		SET status = ?,
		updated_at = datetime('now')
		WHERE email = ?
		`,
		status,
		email,
	)
	return err
}

func UpdateVPNClientStatusByID(id int, status string) error {
	_, err := DB.Exec(
		`
		UPDATE vpn_clients
		SET status = ?,
		updated_at = datetime('now')
		WHERE id = ?
		`,
		status,
		id,
	)
	return err
}

func DeleteVPNClientByID(id int) error {
	_, err := DB.Exec(
		`
		DELETE FROM vpn_clients
		WHERE id = ?
		`,
		id,
	)
	return err
}

func GetAllVPNClients() ([]models.VPNClient, error) {
	rows, err := DB.Query(`SELECT id, email, client_ip, status, created_at FROM vpn_clients`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var clients []models.VPNClient
	for rows.Next() {
		var c models.VPNClient
		err := rows.Scan(&c.ID, &c.Email, &c.ClientIP, &c.Status, &c.CreatedAt)
		if err != nil {
			return nil, err
		}
		clients = append(clients, c)
	}
	return clients, nil
}
