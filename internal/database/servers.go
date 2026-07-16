package database

import (
	"korzadivpn/internal/models"
)

func CreateServers() error {

	servers := []models.Server{

		{
			ID:            1,
			Name:          "Miami-01",
			Country:       "Estados Unidos",
			City:          "Miami",
			Protocol:      "WireGuard",
			Status:        "online",
			MaxUsers:      500,
			CurrentUsers:  120,
			Latency:       18,
			ServerIP:      "miami.korzadivpn.com",
			WireGuardPort: 51820,
			DNS:           "1.1.1.1",
		},

		{
			ID:            2,
			Name:          "Madrid-01",
			Country:       "España",
			City:          "Madrid",
			Protocol:      "WireGuard",
			Status:        "online",
			MaxUsers:      500,
			CurrentUsers:  85,
			Latency:       42,
			ServerIP:      "madrid.korzadivpn.com",
			WireGuardPort: 51820,
			DNS:           "1.1.1.1",
		},

		{
			ID:            3,
			Name:          "São Paulo-01",
			Country:       "Brasil",
			City:          "São Paulo",
			Protocol:      "OpenVPN",
			Status:        "online",
			MaxUsers:      500,
			CurrentUsers:  210,
			Latency:       35,
			ServerIP:      "saopaulo.korzadivpn.com",
			WireGuardPort: 1194,
			DNS:           "1.1.1.1",
		},
	}

	for _, server := range servers {

		_, err := DB.Exec(`
		INSERT OR IGNORE INTO servers
		(
			id,
			name,
			country,
			city,
			protocol,
			status,
			max_users,
			current_users,
			latency,
			server_ip,
			wireguard_port,
			dns
		)
		VALUES(?,?,?,?,?,?,?,?,?,?,?,?)
		`,
			server.ID,
			server.Name,
			server.Country,
			server.City,
			server.Protocol,
			server.Status,
			server.MaxUsers,
			server.CurrentUsers,
			server.Latency,
			server.ServerIP,
			server.WireGuardPort,
			server.DNS,
		)

		if err != nil {
			return err
		}
	}

	return nil
}

func GetServers() ([]models.Server, error) {

	rows, err := DB.Query(`
	SELECT
	id,
	name,
	country,
	city,
	protocol,
	status,
	max_users,
	current_users,
	latency,
	server_ip,
	wireguard_port,
	server_public_key,
	server_private_key,
	dns
	FROM servers
	`)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var servers []models.Server

	for rows.Next() {

		var server models.Server

		err := rows.Scan(
			&server.ID,
			&server.Name,
			&server.Country,
			&server.City,
			&server.Protocol,
			&server.Status,
			&server.MaxUsers,
			&server.CurrentUsers,
			&server.Latency,
			&server.ServerIP,
			&server.WireGuardPort,
			&server.ServerPublicKey,
			&server.ServerPrivateKey,
			&server.DNS,
		)

		if err != nil {
			return nil, err
		}

		servers = append(servers, server)
	}

	return servers, nil
}

func GetServerByID(id int) (*models.Server, error) {

	row := DB.QueryRow(`
	SELECT
	id,
	name,
	country,
	city,
	protocol,
	status,
	max_users,
	current_users,
	latency,
	server_ip,
	wireguard_port,
	server_public_key,
	server_private_key,
	dns
	FROM servers
	WHERE id=?
	`, id)

	var server models.Server

	err := row.Scan(
		&server.ID,
		&server.Name,
		&server.Country,
		&server.City,
		&server.Protocol,
		&server.Status,
		&server.MaxUsers,
		&server.CurrentUsers,
		&server.Latency,
		&server.ServerIP,
		&server.WireGuardPort,
		&server.ServerPublicKey,
		&server.ServerPrivateKey,
		&server.DNS,
	)

	if err != nil {
		return nil, err
	}

	return &server, nil
}

func IncrementServerUsers(id int) error {

	_, err := DB.Exec(`
	UPDATE servers
	SET current_users=current_users+1
	WHERE id=?
	`, id)

	return err
}

func DecrementServerUsers(id int) error {

	_, err := DB.Exec(`
	UPDATE servers
	SET current_users =
	CASE
	WHEN current_users>0 THEN current_users-1
	ELSE 0
	END
	WHERE id=?
	`, id)

	return err
}

func UpdateServerStatus(id int, status string) error {

	_, err := DB.Exec(`
	UPDATE servers
	SET status=?
	WHERE id=?
	`, status, id)

	return err
}

// GetBestServer obtiene el mejor servidor disponible.
func GetBestServer() (*models.Server, error) {

	row := DB.QueryRow(`
	SELECT
	id,
	name,
	country,
	city,
	protocol,
	status,
	max_users,
	current_users,
	latency,
	server_ip,
	wireguard_port,
	server_public_key,
	server_private_key,
	dns
	FROM servers
	WHERE status='online'
	AND current_users < max_users
	ORDER BY
	(latency + (current_users * 100 / max_users)) ASC
	LIMIT 1
	`)

	var server models.Server

	err := row.Scan(

		&server.ID,

		&server.Name,

		&server.Country,

		&server.City,

		&server.Protocol,

		&server.Status,

		&server.MaxUsers,

		&server.CurrentUsers,

		&server.Latency,

		&server.ServerIP,

		&server.WireGuardPort,

		&server.ServerPublicKey,

		&server.ServerPrivateKey,

		&server.DNS,
	)

	if err != nil {

		return nil, err

	}

	return &server, nil

}
