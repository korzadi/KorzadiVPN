package vpncore

type VPNClient struct {
	Email string

	IP string

	PublicKey string

	PrivateKey string
}

func CreateClient(
	email string,
) VPNClient {

	return VPNClient{

		Email: email,

		IP: "10.8.0.2",
	}
}
