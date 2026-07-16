package models

type Connection struct {
	ID int `json:"id"`

	Email string `json:"email"`

	ServerID int `json:"server_id"`

	Server string `json:"server"`

	Status string `json:"status"`

	Device string `json:"device"`

	ClientID string `json:"client_id"`

	IP string `json:"ip"`

	ConnectedAt string `json:"connected_at"`

	DisconnectedAt string `json:"disconnected_at"`

	LastPing string `json:"last_ping"`
}

var Connections []Connection
