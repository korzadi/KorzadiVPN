package monitor

type ServerStatus struct {
	Server string

	Online bool

	Load int
}

func Check(server string) ServerStatus {

	return ServerStatus{

		Server: server,

		Online: true,

		Load: 0,
	}
}
