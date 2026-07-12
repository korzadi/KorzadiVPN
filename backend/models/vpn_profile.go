package models

type VPNProfile struct {
	Email      string `json:"email"`
	ServerID   int    `json:"server_id"`
	Server     string `json:"server"`
	Protocol   string `json:"protocol"`
	PublicKey  string `json:"public_key"`
	PrivateKey string `json:"private_key"`
	Config     string `json:"config"`
	Status     string `json:"status"`
}

var VPNProfiles []VPNProfile
