package wireguard

type WireGuardManager interface {
	AddPeer(publicKey string, ip string) error
	RemovePeer(publicKey string) error
}
