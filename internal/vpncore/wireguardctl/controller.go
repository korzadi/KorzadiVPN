package wireguardctl

func CreateInterface() error {
	return RunSafe("wg", "show")
}

func AddPeer(
	publicKey string,
	ip string,
) error {
	return RunSafe("wg", "set", "wg0", "peer", publicKey, "allowed-ips", ip+"/32")
}

func RemovePeer(
	publicKey string,
) error {
	return RunSafe("wg", "set", "wg0", "peer", publicKey, "remove")
}
