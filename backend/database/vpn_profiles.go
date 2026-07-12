package database

import (
	"korzadivpn/models"
)

func CreateVPNProfile(profile models.VPNProfile) error {

	_, err := DB.Exec(`
	INSERT INTO vpn_profiles
	(
		email,
		server_id,
		server,
		protocol,
		public_key,
		private_key,
		config,
		status
	)
	VALUES (?,?,?,?,?,?,?,?)
	`,
		profile.Email,
		profile.ServerID,
		profile.Server,
		profile.Protocol,
		profile.PublicKey,
		profile.PrivateKey,
		profile.Config,
		profile.Status,
	)

	return err
}

func GetVPNProfile(email string) (*models.VPNProfile, error) {

	row := DB.QueryRow(`
	SELECT
		email,
		server_id,
		server,
		protocol,
		public_key,
		private_key,
		config,
		status
	FROM vpn_profiles
	WHERE email = ?
	LIMIT 1
	`,
		email,
	)

	var profile models.VPNProfile

	err := row.Scan(
		&profile.Email,
		&profile.ServerID,
		&profile.Server,
		&profile.Protocol,
		&profile.PublicKey,
		&profile.PrivateKey,
		&profile.Config,
		&profile.Status,
	)

	if err != nil {
		return nil, err
	}

	return &profile, nil
}
