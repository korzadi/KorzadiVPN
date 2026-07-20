package database

import (
	"korzadivpn/internal/models"
)

func CreateVPNProfile(profile models.VPNProfile) error {

	// Verificar si ya existe un perfil para este usuario y servidor
	var exists int

	err := DB.QueryRow(`
                SELECT COUNT(*)
                FROM vpn_profiles
                WHERE email = ? AND server_id = ?
        `,
		profile.Email,
		profile.ServerID,
	).Scan(&exists)

	if err != nil {
		return err
	}

	// Si existe, actualizarlo
	if exists > 0 {

		_, err := DB.Exec(`
                        UPDATE vpn_profiles
                        SET
                        server = ?,
                        protocol = ?,
                        public_key = ?,
                        private_key = ?,
                        config = ?,
                        status = ?
                        WHERE email = ? AND server_id = ?
                `,
			profile.Server,
			profile.Protocol,
			profile.PublicKey,
			profile.PrivateKey,
			profile.Config,
			profile.Status,
			profile.Email,
			profile.ServerID,
		)

		return err
	}

	// Si no existe, crear uno nuevo
	_, err = DB.Exec(`
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
