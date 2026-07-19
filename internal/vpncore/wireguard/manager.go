package wireguard

type Manager struct {
	Server    Server
	Interface InterfaceController
	Peers     PeerController
	Firewall  FirewallManager
}

func NewManager(
	server Server,
) Manager {

	return Manager{

		Server: server,

		Interface: NewInterfaceController(),

		Peers: NewPeerController(),

		Firewall: NewFirewallManager(),
	}
}

func (m Manager) ServerConfig() string {

	return NewServerManager(
		m.Server,
	).GenerateInterfaceConfig()
}

func (m Manager) AddClient(
	publicKey string,
	ip string,
) string {

	return m.Peers.AddPeer(
		publicKey,
		ip,
	)
}

func (m Manager) RemoveClient(
	publicKey string,
) string {

	return m.Peers.RemovePeer(
		publicKey,
	)
}
