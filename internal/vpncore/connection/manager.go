package connection

import (
	"korzadivpn/internal/models"
	"korzadivpn/internal/vpncore/wireguard"
)

type ConnectionManager struct {
	WireGuard wireguard.Manager
}

func NewConnectionManager(
	manager wireguard.Manager,
) ConnectionManager {

	return ConnectionManager{
		WireGuard: manager,
	}
}

func (c ConnectionManager) ConnectClient(
	client models.VPNClient,
) string {

	return c.WireGuard.AddClient(
		client.PublicKey,
		client.ClientIP,
	)
}

func (c ConnectionManager) DisconnectClient(
	client models.VPNClient,
) string {

	return c.WireGuard.RemoveClient(
		client.PublicKey,
	)
}
