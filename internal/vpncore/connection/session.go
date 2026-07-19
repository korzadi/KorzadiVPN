package connection

import "time"

type VPNConnection struct {
	Email string

	Server string

	IP string

	Status string

	ConnectedAt string
}

func Start(
	email string,
	server string,
	ip string,
) VPNConnection {

	return VPNConnection{

		Email: email,

		Server: server,

		IP: ip,

		Status: "connected",

		ConnectedAt: time.Now().
			UTC().
			Format(time.RFC3339),
	}
}

func Stop(
	conn VPNConnection,
) VPNConnection {

	conn.Status = "disconnected"

	return conn
}
