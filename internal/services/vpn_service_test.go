package services

import (
	"encoding/base64"
	"errors"
	"korzadivpn/internal/models"
	"korzadivpn/internal/vpncore/wireguard"
	"testing"
)

type mockDB struct {
	UpdateVPNConnectionStatusFunc func(email string, status string) error
}

func (m *mockDB) UpdateVPNConnectionStatus(email string, status string) error {
	return m.UpdateVPNConnectionStatusFunc(email, status)
}

func TestProvisionPeer(t *testing.T) {
	mockWG := &wireguard.MockWireGuardManager{}
	mockDB := &mockDB{}
	service := VPNService{WireGuard: mockWG, DB: mockDB}

	t.Run("Success", func(t *testing.T) {
		mockWG.AddPeerFunc = func(pk string, ip string) error {
			return nil
		}

		client := &models.VPNClient{
			PublicKey: "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuv",
			ClientIP:  "10.10.0.2",
		}

		err := service.ProvisionPeer(client)
		if err != nil {
			t.Errorf("Expected success, got %v", err)
		}
		if mockWG.AddPeerCalled != 1 {
			t.Errorf("Expected 1 call, got %d", mockWG.AddPeerCalled)
		}
	})

	t.Run("Failure", func(t *testing.T) {
		mockWG.AddPeerFunc = func(pk string, ip string) error {
			return errors.New("wg failed")
		}

		err := service.ProvisionPeer(&models.VPNClient{})
		if err == nil {
			t.Error("Expected error, got nil")
		}
	})
}

func TestDisconnect(t *testing.T) {
	mockWG := &wireguard.MockWireGuardManager{}
	mockDB := &mockDB{
		UpdateVPNConnectionStatusFunc: func(email string, status string) error {
			return nil
		},
	}
	service := VPNService{WireGuard: mockWG, DB: mockDB}

	t.Run("Success", func(t *testing.T) {
		mockWG.RemovePeerFunc = func(pk string) error {
			return nil
		}

		err := service.Disconnect(&models.VPNClient{PublicKey: "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuv"})
		if err != nil {
			t.Errorf("Expected success, got %v", err)
		}
		if mockWG.RemovePeerCalled != 1 {
			t.Errorf("Expected 1 call, got %d", mockWG.RemovePeerCalled)
		}
	})
}
func TestValidators(t *testing.T) {
	t.Run("PublicKey", func(t *testing.T) {
		// Generar una clave válida de 32 bytes en base64
		key := make([]byte, 32)
		validKey := base64.StdEncoding.EncodeToString(key)
		if !wireguard.IsValidPublicKey(validKey) {
			t.Errorf("Valid key failed: %s (len: %d)", validKey, len(validKey))
		}

		if wireguard.IsValidPublicKey("short") {
			t.Error("Invalid key passed")
		}
	})
	t.Run("IP", func(t *testing.T) {
		if !wireguard.IsValidCIDR("10.10.0.2/32") {
			t.Error("Valid CIDR failed")
		}
		if wireguard.IsValidCIDR("invalid") {
			t.Error("Invalid CIDR passed")
		}
	})
}
