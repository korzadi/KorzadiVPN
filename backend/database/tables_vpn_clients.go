package database

import "log"

func CreateVPNClientTable() {

	query := `
	CREATE TABLE IF NOT EXISTS vpn_clients (

		id INTEGER PRIMARY KEY AUTOINCREMENT,

		email TEXT NOT NULL,

		server_id INTEGER,

		node_id INTEGER,

		client_name TEXT NOT NULL,

		device_id TEXT,

		device_name TEXT,

		device_type TEXT,

		client_ip TEXT,

		ipv6 TEXT,

		public_key TEXT,

		private_key TEXT,

		preshared_key TEXT,

		config TEXT,

		protocol TEXT,

		dns TEXT,

		mtu INTEGER,

		allowed_ips TEXT,

		endpoint TEXT,

		status TEXT,

		connection_status TEXT,

		plan TEXT,

		bandwidth_limit INTEGER,

		data_used INTEGER,

		max_devices INTEGER,

		last_handshake TEXT,

		last_connected TEXT,

		last_disconnected TEXT,

		last_ip TEXT,

		country TEXT,

		city TEXT,

		expires_at TEXT,

		revoked_at TEXT,

		created_at TEXT,

		updated_at TEXT

	);
	`

	_, err := DB.Exec(query)

	if err != nil {
		log.Fatal(err)
	}

	log.Println("Tabla vpn_clients creada correctamente")
}
