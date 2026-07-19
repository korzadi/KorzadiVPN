package wireguard

type Peer struct {
	Email      string
	PublicKey  string
	PrivateKey string
	IPAddress  string
}

type Server struct {
	Name      string
	PublicKey string
	Endpoint  string
	Port      int
}
