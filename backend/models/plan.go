package models

type Plan struct {
	Name        string `json:"name"`
	MaxDevices  int    `json:"max_devices"`
	Description string `json:"description"`
}
