package models

type VPNProfile struct {
	Email string `json:"email"`

	ServerID int `json:"server_id"`

	Server string `json:"server"`

	Protocol string `json:"protocol"`

	PublicKey string `json:"public_key"`

	// Se usa internamente para generar el archivo WireGuard
	// pero nunca se envía directamente en JSON
	PrivateKey string `json:"-"`

	Config string `json:"config"`

	Status string `json:"status"`
}

var VPNProfiles []VPNProfile
