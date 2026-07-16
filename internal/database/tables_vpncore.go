package database

import "log"

func CreateVPNCoreTables() {

	query := `
	CREATE TABLE IF NOT EXISTS vpn_core_clients (

		id INTEGER PRIMARY KEY AUTOINCREMENT,

		email TEXT,

		ip TEXT,

		public_key TEXT,

		private_key TEXT,

		status TEXT,

		created_at TEXT

	);
	`

	_, err := DB.Exec(query)

	if err != nil {
		log.Fatal(err)
	}

	log.Println("Tabla vpn_core_clients creada correctamente")
}
