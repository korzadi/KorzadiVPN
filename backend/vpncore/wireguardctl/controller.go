package wireguardctl

import (
	"fmt"
	"os/exec"
)

func CreateInterface() error {

	cmd := exec.Command(
		"wg",
		"show",
	)

	_, err := cmd.Output()

	if err != nil {
		fmt.Println(
			"WireGuard no disponible todavía:",
			err,
		)
	}

	return nil
}

func AddPeer(
	publicKey string,
	ip string,
) error {

	cmd := exec.Command(
		"wg",
		"set",
		"wg0",
		"peer",
		publicKey,
		"allowed-ips",
		ip,
	)

	return cmd.Run()
}

func RemovePeer(
	publicKey string,
) error {

	cmd := exec.Command(
		"wg",
		"set",
		"wg0",
		"peer",
		publicKey,
		"remove",
	)

	return cmd.Run()
}
