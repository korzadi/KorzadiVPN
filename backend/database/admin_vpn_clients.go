package database

import (
	"korzadivpn/models"
)

func GetAllVPNClients() ([]models.VPNClient, error) {

	rows, err := DB.Query(`
	SELECT
	id,
	email,
	server_id,
	client_name,
	client_ip,
	public_key,
	status,
	created_at
	FROM vpn_clients
	`)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var clients []models.VPNClient

	for rows.Next() {

		var client models.VPNClient

		err := rows.Scan(
			&client.ID,
			&client.Email,
			&client.ServerID,
			&client.ClientName,
			&client.ClientIP,
			&client.PublicKey,
			&client.Status,
			&client.CreatedAt,
		)

		if err != nil {
			return nil, err
		}

		clients = append(
			clients,
			client,
		)
	}

	return clients, nil
}

func UpdateVPNClientStatus(
	id int,
	status string,
) error {

	_, err := DB.Exec(
		`
		UPDATE vpn_clients
		SET status=?
		WHERE id=?
		`,
		status,
		id,
	)

	return err
}

func DeleteVPNClientByID(
	id int,
) error {

	_, err := DB.Exec(
		`
		DELETE FROM vpn_clients
		WHERE id=?
		`,
		id,
	)

	return err
}
