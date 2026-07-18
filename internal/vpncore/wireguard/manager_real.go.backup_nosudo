package wireguard

import (
	"fmt"
	"log"
	"os/exec"
)

var _ WireGuardManager = (*RealManager)(nil)

type RealManager struct {
	Interface string
}

func NewRealManager() RealManager {
	return RealManager{
		Interface: "wg0",
	}
}

// AddPeer añade un peer a la interfaz WireGuard mediante el comando sudo wg set
func (m RealManager) AddPeer(publicKey string, ip string) error {
	// sudo wg set wg0 peer <pubkey> allowed-ips <ip>/32
	args := []string{"wg", "set", m.Interface, "peer", publicKey, "allowed-ips", ip + "/32"}
	cmd := exec.Command("sudo", args...)

	log.Printf("Aprovisionando peer en %s (root): %v", m.Interface, args)

	output, err := cmd.CombinedOutput()
	if err != nil {
		return fmt.Errorf("error al aprovisionar peer %s (ejecutar 'visudo' para dar permisos sudo): %v, salida: %s", publicKey, err, string(output))
	}

	return nil
}

// RemovePeer elimina un peer de la interfaz WireGuard mediante el comando sudo wg set
func (m RealManager) RemovePeer(publicKey string) error {
	// sudo wg set wg0 peer <pubkey> remove
	args := []string{"wg", "set", m.Interface, "peer", publicKey, "remove"}
	cmd := exec.Command("sudo", args...)

	log.Printf("Eliminando peer de %s (root): %v", m.Interface, args)

	output, err := cmd.CombinedOutput()
	if err != nil {
		return fmt.Errorf("error al eliminar peer %s: %v, salida: %s", publicKey, err, string(output))
	}

	return nil
}

// Status devuelve el estado actual de la interfaz WireGuard
func (m RealManager) Status() ([]byte, error) {
	cmd := exec.Command("wg", "show", m.Interface)
	return cmd.Output()
}
