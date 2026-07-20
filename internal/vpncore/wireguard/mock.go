package wireguard

type MockWireGuardManager struct {
	AddPeerFunc      func(publicKey string, ip string) error
	RemovePeerFunc   func(publicKey string) error
	AddPeerCalled    int
	RemovePeerCalled int
}

func (m *MockWireGuardManager) AddPeer(publicKey string, ip string) error {
	m.AddPeerCalled++
	if m.AddPeerFunc != nil {
		return m.AddPeerFunc(publicKey, ip)
	}
	return nil
}

func (m *MockWireGuardManager) RemovePeer(publicKey string) error {
	m.RemovePeerCalled++
	if m.RemovePeerFunc != nil {
		return m.RemovePeerFunc(publicKey)
	}
	return nil
}
