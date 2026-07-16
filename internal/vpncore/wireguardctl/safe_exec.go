package wireguardctl

import (
	"fmt"
	"korzadivpn/internal/vpncore/wireguard"
	"os/exec"
)

// RunSafe ejecutan comandos de wireguard validando los argumentos
func RunSafe(cmdName string, args ...string) error {
	// Validación de seguridad básica
	if cmdName != "wg" {
		return fmt.Errorf("comando no permitido: %s", cmdName)
	}

	// Validación de argumentos según el subcomando
	if len(args) > 0 {
		subcommand := args[0]
		switch subcommand {
		case "set":
			// wg set wg0 peer <pk> allowed-ips <ip>/32
			if len(args) < 6 {
				return fmt.Errorf("argumentos insuficientes para set")
			}
			pk := args[3]
			cidr := args[5]
			if !wireguard.IsValidPublicKey(pk) {
				return fmt.Errorf("public key invalida")
			}
			if !wireguard.IsValidCIDR(cidr) {
				return fmt.Errorf("CIDR invalido")
			}
		case "show":
			// Permitido
		default:
			return fmt.Errorf("subcomando no soportado: %s", subcommand)
		}
	}

	cmd := exec.Command(cmdName, args...)
	return cmd.Run()
}
