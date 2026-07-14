package database

import (
	"korzadivpn/utils"
)

func MigrateServerManagement() error {

	_, _ = DB.Exec(`
	ALTER TABLE servers
	ADD COLUMN server_public_key TEXT;
	`)

	_, _ = DB.Exec(`
	ALTER TABLE servers
	ADD COLUMN server_private_key TEXT;
	`)

	rows, err := DB.Query(`
	SELECT id
	FROM servers
	WHERE server_public_key IS NULL
	OR server_private_key IS NULL
	`)

	if err != nil {
		return err
	}

	var ids []int

	for rows.Next() {

		var id int

		err := rows.Scan(&id)

		if err != nil {
			continue
		}

		ids = append(ids, id)
	}

	rows.Close()

	for _, id := range ids {

		publicKey := utils.GenerateServerKey()

		privateKey := utils.GenerateServerKey()

		_, err := DB.Exec(`
		UPDATE servers
		SET
		server_public_key=?,
		server_private_key=?
		WHERE id=?
		`,
			publicKey,
			privateKey,
			id,
		)

		if err != nil {
			return err
		}
	}

	return nil
}
