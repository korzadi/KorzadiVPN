package database

import (
	"fmt"
)

// GetNextVPNClientIP genera una IP libre para un cliente WireGuard.
func GetNextVPNClientIP() (string, error) {

	for i := 2; i < 255; i++ {

		ip := fmt.Sprintf(
			"10.0.0.%d",
			i,
		)

		var count int

		err := DB.QueryRow(
			`
			SELECT COUNT(*)
			FROM vpn_clients
			WHERE client_ip=?
			`,
			ip,
		).Scan(
			&count,
		)

		if err != nil {
			return "", err
		}

		if count == 0 {

			return ip, nil

		}

	}

	return "", fmt.Errorf(
		"No hay IPs disponibles",
	)

}
