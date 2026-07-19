package config

type NodeConfig struct {
	ID string

	Name string

	Host string

	Port int

	Region string

	Status string

	MaxClients int

	CurrentClients int
}

func NewNode(
	id string,
	name string,
	host string,
	region string,
) NodeConfig {

	return NodeConfig{

		ID: id,

		Name: name,

		Host: host,

		Port: 51820,

		Region: region,

		Status: "offline",

		MaxClients: 1000,

		CurrentClients: 0,
	}
}
