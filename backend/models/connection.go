package models

type Connection struct {
	Email string `json:"email"`

	ServerID int `json:"server_id"`

	Server string `json:"server"`

	Status string `json:"status"`

	Device string `json:"device"`

	IP string `json:"ip"`

	ConnectedAt string `json:"connected_at"`
}

var Connections []Connection
