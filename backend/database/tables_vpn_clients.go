package database

import "log"

func CreateVPNClientTable() {

	query := `
	CREATE TABLE IF NOT EXISTS vpn_clients (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		email TEXT NOT NULL,
		server_id INTEGER,
		client_name TEXT NOT NULL,
		client_ip TEXT,
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

	log.Println("Tabla vpn_clients creada correctamente")
}
