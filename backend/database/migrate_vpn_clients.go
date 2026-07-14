package database

import (
	"log"
)

func columnExists(
	table string,
	column string,
) bool {

	var name string

	err := DB.QueryRow(
		"PRAGMA table_info(" + table + ")",
	).Scan(
		&name,
	)

	if err != nil {
		return false
	}

	return name == column
}

func addColumn(
	column string,
	definition string,
) {

	var count int

	err := DB.QueryRow(
		`
		SELECT COUNT(*)
		FROM pragma_table_info('vpn_clients')
		WHERE name=?
		`,
		column,
	).Scan(&count)

	if err != nil {
		return
	}

	if count == 0 {

		_, err :=
			DB.Exec(
				"ALTER TABLE vpn_clients ADD COLUMN " +
					column + " " + definition,
			)

		if err != nil {
			log.Println(
				"Error agregando columna",
				column,
				err,
			)
		}
	}
}

func MigrateVPNClients() {

	columns := map[string]string{

		"node_id":           "INTEGER",
		"device_id":         "TEXT",
		"device_name":       "TEXT",
		"device_type":       "TEXT",
		"ipv6":              "TEXT",
		"preshared_key":     "TEXT",
		"config":            "TEXT",
		"protocol":          "TEXT",
		"dns":               "TEXT",
		"mtu":               "INTEGER",
		"allowed_ips":       "TEXT",
		"endpoint":          "TEXT",
		"connection_status": "TEXT",
		"plan":              "TEXT",
		"bandwidth_limit":   "INTEGER",
		"data_used":         "INTEGER",
		"max_devices":       "INTEGER",
		"last_handshake":    "TEXT",
		"last_connected":    "TEXT",
		"last_disconnected": "TEXT",
		"last_ip":           "TEXT",
		"country":           "TEXT",
		"city":              "TEXT",
		"expires_at":        "TEXT",
		"revoked_at":        "TEXT",
		"updated_at":        "TEXT",
	}

	for column, definition := range columns {

		addColumn(
			column,
			definition,
		)

	}

	log.Println(
		"Migración vpn_clients completada",
	)

}
