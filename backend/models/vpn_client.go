package models

type VPNClient struct {
	ID int `json:"id"`

	Email string `json:"email"`

	ServerID int `json:"server_id"`

	NodeID int `json:"node_id"`

	ClientName string `json:"client_name"`

	DeviceID string `json:"device_id"`

	DeviceName string `json:"device_name"`

	DeviceType string `json:"device_type"`

	ClientIP string `json:"client_ip"`

	IPv6 string `json:"ipv6"`

	PublicKey string `json:"public_key"`

	PrivateKey string `json:"private_key"`

	PresharedKey string `json:"preshared_key"`

	Config string `json:"config"`

	Protocol string `json:"protocol"`

	DNS string `json:"dns"`

	MTU int `json:"mtu"`

	AllowedIPs string `json:"allowed_ips"`

	Endpoint string `json:"endpoint"`

	Status string `json:"status"`

	ConnectionStatus string `json:"connection_status"`

	Plan string `json:"plan"`

	BandwidthLimit int64 `json:"bandwidth_limit"`

	DataUsed int64 `json:"data_used"`

	MaxDevices int `json:"max_devices"`

	LastHandshake string `json:"last_handshake"`

	LastConnected string `json:"last_connected"`

	LastDisconnected string `json:"last_disconnected"`

	LastIP string `json:"last_ip"`

	Country string `json:"country"`

	City string `json:"city"`

	ExpiresAt string `json:"expires_at"`

	RevokedAt string `json:"revoked_at"`

	CreatedAt string `json:"created_at"`

	UpdatedAt string `json:"updated_at"`
}
