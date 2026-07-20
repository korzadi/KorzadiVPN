package database

import "log"

func MigrateServerKeys() {

	var count int

	err := DB.QueryRow(`
	SELECT COUNT(*)
	FROM pragma_table_info('servers')
	WHERE name='server_private_key'
	`).Scan(&count)

	if err != nil {
		log.Println("Error revisando migración:", err)
		return
	}

	if count == 0 {

		_, err := DB.Exec(`
		ALTER TABLE servers
		ADD COLUMN server_private_key TEXT;
		`)

		if err != nil {
			log.Println("Error agregando server_private_key:", err)
			return
		}

		log.Println("Columna server_private_key creada")
		return
	}

	log.Println("Migración server_private_key: OK")
}
