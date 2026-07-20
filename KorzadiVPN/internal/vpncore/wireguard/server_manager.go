package wireguard

import (
	"fmt"
)

type ServerManager struct {
	Interface string
	Server    Server
}

func NewServerManager(
	server Server,
) ServerManager {

	return ServerManager{

		Interface: "wg0",

		Server: server,
	}
}

func (m ServerManager) GenerateInterfaceConfig() string {

	return fmt.Sprintf(
		`[Interface]
PrivateKey = %s
Address = %s
ListenPort = %d
`,
		m.Server.PrivateKey,
		m.Server.Address,
		m.Server.Port,
	)
}
