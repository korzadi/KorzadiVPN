package models

var Servers = []Server{
	{
		ID:           1,
		Name:         "Miami-01",
		Country:      "Estados Unidos",
		City:         "Miami",
		Protocol:     "WireGuard",
		Status:       "online",
		MaxUsers:     500,
		CurrentUsers: 120,
		Latency:      18,
	},
	{
		ID:           2,
		Name:         "Madrid-01",
		Country:      "España",
		City:         "Madrid",
		Protocol:     "WireGuard",
		Status:       "online",
		MaxUsers:     500,
		CurrentUsers: 85,
		Latency:      42,
	},
	{
		ID:           3,
		Name:         "São Paulo-01",
		Country:      "Brasil",
		City:         "São Paulo",
		Protocol:     "OpenVPN",
		Status:       "online",
		MaxUsers:     500,
		CurrentUsers: 210,
		Latency:      35,
	},
}
