package models

type AppConfig struct {
	Version string `json:"version"`

	API string `json:"api"`

	VPNCore string `json:"vpn_core"`

	Status string `json:"status"`
}
