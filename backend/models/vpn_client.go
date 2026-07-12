package models

type VPNClient struct {

	ID int `json:"id"`

	Email string `json:"email"`

	ServerID int `json:"server_id"`

	ClientName string `json:"client_name"`

	ClientIP string `json:"client_ip"`

	PublicKey string `json:"public_key"`

	PrivateKey string `json:"private_key"`

	Status string `json:"status"`

	CreatedAt string `json:"created_at"`
}
