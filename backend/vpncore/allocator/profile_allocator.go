package allocator

import (
	"korzadivpn/vpncore/profile"
)

func CreateUserVPNProfile(
	email string,
	server string,
	ip string,
	privateKey string,
	publicKey string,
) profile.VPNProfile {

	config := profile.GenerateWireGuardConfig(
		privateKey,
		ip,
		"1.1.1.1",
		publicKey,
		server,
	)

	return profile.NewProfile(
		email,
		server,
		ip,
		privateKey,
		publicKey,
		config,
	)
}
