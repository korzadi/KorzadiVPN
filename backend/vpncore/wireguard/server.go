package wireguard

type Server struct {
	Name string

	PublicKey string

	PrivateKey string

	Endpoint string

	Port int

	Address string
}

func NewServer(
	name string,
	endpoint string,
	port int,
) Server {

	return Server{

		Name: name,

		Endpoint: endpoint,

		Port: port,

		Address: "10.8.0.1/24",

		PublicKey: GenerateKey(),

		PrivateKey: GenerateKey(),
	}
}

func (s Server) Config() string {

	return ""

}
