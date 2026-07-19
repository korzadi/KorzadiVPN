package vpncore

type ClientConfig struct {
	Email string

	Server string

	PublicKey string

	PrivateKey string
}

func CreateClientConfig(
	email string,
	server string,
	ip string,
) string {

	return "VPN CONFIG\n" +
		"EMAIL=" + email + "\n" +
		"SERVER=" + server + "\n" +
		"IP=" + ip
}
