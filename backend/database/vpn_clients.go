package database

import (
	"korzadivpn/models"
)

func CreateVPNClient(client models.VPNClient) error {

	_, err := DB.Exec(
		`
		INSERT INTO vpn_clients
		(
			email,
			server_id,
			client_name,
			client_ip,
			public_key,
			private_key,
			status,
			created_at
		)
		VALUES (?,?,?,?,?,?,?,?)
		`,
		client.Email,
		client.ServerID,
		client.ClientName,
		client.ClientIP,
		client.PublicKey,
		client.PrivateKey,
		client.Status,
		client.CreatedAt,
	)

	return err
}


func GetVPNClientsByEmail(email string) ([]models.VPNClient, error) {

	rows, err := DB.Query(
		`
		SELECT
			id,
			email,
			server_id,
			client_name,
			client_ip,
			public_key,
			private_key,
			status,
			created_at
		FROM vpn_clients
		WHERE email=?
		`,
		email,
	)

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
			&client.PrivateKey,
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


func GetVPNClientByEmail(email string) (*models.VPNClient, error) {

        row := DB.QueryRow(
                `
                SELECT
                        id,
                        email,
                        server_id,
                        client_name,
                        client_ip,
                        public_key,
                        private_key,
                        status,
                        created_at
                FROM vpn_clients
                WHERE email=?
                LIMIT 1
                `,
                email,
        )

        var client models.VPNClient

        err := row.Scan(
                &client.ID,
                &client.Email,
                &client.ServerID,
                &client.ClientName,
                &client.ClientIP,
                &client.PublicKey,
                &client.PrivateKey,
                &client.Status,
                &client.CreatedAt,
        )

        if err != nil {
                return nil, err
        }

        return &client, nil
}
