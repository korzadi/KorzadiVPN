package services

import (
	"korzadivpn/internal/database"
	"korzadivpn/internal/models"
	"korzadivpn/internal/vpncore/wireguard"
)

type VPNService struct {
	WireGuard wireguard.WireGuardManager
	DB        DBInterface
}

type DBInterface interface {
	UpdateVPNConnectionStatus(email string, status string) error
}

type realDB struct{}

func (realDB) UpdateVPNConnectionStatus(email string, status string) error {
	return database.UpdateVPNConnectionStatus(email, status)
}

func NewVPNService() VPNService {
	return VPNService{
		WireGuard: wireguard.NewRealManager(),
		DB:        realDB{},
	}
}

func (s VPNService) Connect(
	client *models.VPNClient,
) error {
	err := s.WireGuard.AddPeer(
		client.PublicKey,
		client.ClientIP,
	)

	if err != nil {
		return err
	}

	return s.DB.UpdateVPNConnectionStatus(
		client.Email,
		"connected",
	)
}

func (s VPNService) ProvisionPeer(
	client *models.VPNClient,
) error {
	return s.WireGuard.AddPeer(client.PublicKey, client.ClientIP)
}

func (s VPNService) Disconnect(
	client *models.VPNClient,
) error {

	err := s.WireGuard.RemovePeer(
		client.PublicKey,
	)

	if err != nil {
		return err
	}

	return s.DB.UpdateVPNConnectionStatus(
		client.Email,
		"disconnected",
	)
}
