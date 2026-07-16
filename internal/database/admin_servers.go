package database

import (
	"errors"

	"korzadivpn/internal/models"
	"korzadivpn/pkg/utils"
)

func CreateServer(server models.Server) error {

	if server.ServerPublicKey == "" {
		server.ServerPublicKey = utils.GenerateServerKey()
	}

	if server.ServerPrivateKey == "" {
		server.ServerPrivateKey = utils.GenerateServerKey()
	}

	_, err := DB.Exec(`
	INSERT INTO servers
	(
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
	)
	VALUES(?,?,?,?,?,?,?,?,?,?,?,?,?)
	`,
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
		server.ServerPublicKey,
		server.ServerPrivateKey,
		server.DNS,
	)

	return err
}

func UpdateServer(server models.Server) error {

	_, err := DB.Exec(`
	UPDATE servers
	SET
	name=?,
	country=?,
	city=?,
	protocol=?,
	status=?,
	max_users=?,
	latency=?,
	server_ip=?,
	wireguard_port=?,
	dns=?
	WHERE id=?
	`,
		server.Name,
		server.Country,
		server.City,
		server.Protocol,
		server.Status,
		server.MaxUsers,
		server.Latency,
		server.ServerIP,
		server.WireGuardPort,
		server.DNS,
		server.ID,
	)

	return err
}

func GetServerAdminByID(id int) (*models.Server, error) {

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

func DeleteServer(id int) error {

	var users int

	err := DB.QueryRow(`
	SELECT current_users
	FROM servers
	WHERE id=?
	`, id).Scan(&users)

	if err != nil {
		return err
	}

	if users > 0 {

		return errors.New(
			"no se puede eliminar servidor con usuarios conectados",
		)

	}

	_, err = DB.Exec(`
	DELETE FROM servers
	WHERE id=?
	`, id)

	return err
}
