package wireguard

import (
	"encoding/base64"
	"net"
)

// IsValidPublicKey valida que la clave sea base64 de 32 bytes (44 caracteres)
func IsValidPublicKey(pk string) bool {
	if len(pk) != 44 {
		return false
	}
	data, err := base64.StdEncoding.DecodeString(pk)
	if err != nil {
		return false
	}
	return len(data) == 32
}

// IsValidIP valida que sea una IPv4 en formato x.x.x.x
func IsValidIP(ip string) bool {
	parsedIP := net.ParseIP(ip)
	return parsedIP != nil && parsedIP.To4() != nil
}

// IsValidCIDR valida formato ip/mask (e.g., 10.0.0.x/32)
func IsValidCIDR(cidr string) bool {
	ip, _, err := net.ParseCIDR(cidr)
	return err == nil && ip.To4() != nil
}
