package database

import "log"

func CreateTables() {

	users := `
	CREATE TABLE IF NOT EXISTS users (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		email TEXT UNIQUE,
		password TEXT,
		plan TEXT,
		status TEXT
	);
	`

	servers := `
	CREATE TABLE IF NOT EXISTS servers (
		id INTEGER PRIMARY KEY,
		name TEXT,
		country TEXT,
		city TEXT,
		protocol TEXT,
		status TEXT,
		max_users INTEGER,
		current_users INTEGER,
		latency INTEGER,
		server_ip TEXT,
		wireguard_port INTEGER,
		server_public_key TEXT,
		dns TEXT
	);
	`

	connections := `
	CREATE TABLE IF NOT EXISTS connections (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		email TEXT,
		server_id INTEGER,
		server TEXT,
		status TEXT
	);
	`

	vpnProfiles := `
	CREATE TABLE IF NOT EXISTS vpn_profiles (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		email TEXT,
		server_id INTEGER,
		server TEXT,
		protocol TEXT,
		public_key TEXT,
		private_key TEXT,
		config TEXT,
		status TEXT
	);
	`

	sessions := `
	CREATE TABLE IF NOT EXISTS sessions (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		email TEXT,
		token TEXT UNIQUE,
		ip TEXT,
		device TEXT,
		created_at TEXT,
		expires_at TEXT,
		status TEXT
	);
	`

	queries := []string{
		users,
		servers,
		connections,
		vpnProfiles,
		sessions,
	}

	for _, query := range queries {

		_, err := DB.Exec(query)

		if err != nil {
			log.Fatal(err)
		}
	}

	log.Println("Tablas creadas correctamente")
}
