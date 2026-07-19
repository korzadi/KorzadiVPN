package models

type Plan struct {
	Name        string `json:"name"`
	MaxDevices  int    `json:"max_devices"`
	Description string `json:"description"`
}

func GetDeviceLimit(plan string) int {

	for _, p := range Plans {

		if p.Name == plan {

			return p.MaxDevices
		}
	}

	// Seguridad: si el plan no existe,
	// usar límite gratuito.

	return 1
}
