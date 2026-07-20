package wireguard

type Client struct {
	Email      string
	Address    string
	PublicKey  string
	PrivateKey string
	Server     string
	Port       int
	DNS        string
}

func NewClient(email string) Client {
	return Client{
		Email:   email,
		Address: "10.10.0.2/32",
		Port:    51820,
		DNS:     "1.1.1.1",
	}
}
