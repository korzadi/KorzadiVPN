package models

var Plans = []Plan{

	{
		Name:        "free",
		MaxDevices:  1,
		Description: "Plan gratuito de KorzadiVPN",
	},

	{
		Name:        "premium",
		MaxDevices:  5,
		Description: "Plan premium con más conexiones",
	},

	{
		Name:        "enterprise",
		MaxDevices:  20,
		Description: "Plan empresarial",
	},
}
