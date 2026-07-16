package database

func UpdateVPNConnectionStatus(
	email string,
	status string,
) error {

	_, err := DB.Exec(
		`
		UPDATE vpn_clients
		SET
			connection_status=?,
			updated_at=datetime('now')
		WHERE email=?
		`,
		status,
		email,
	)

	return err
}
