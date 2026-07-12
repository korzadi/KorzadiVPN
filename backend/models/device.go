package models

type Device struct {
	ID         int    `json:"id"`
	Email      string `json:"email"`
	DeviceName string `json:"device_name"`
	DeviceType string `json:"device_type"`
	Status     string `json:"status"`
	LastIP     string `json:"last_ip"`
	LastServer string `json:"last_server"`
	LastSeen   string `json:"last_seen"`
	CreatedAt  string `json:"created_at"`
}
