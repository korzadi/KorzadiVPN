package profile

type VPNProfile struct {
	Email string

	Server string

	IP string

	PrivateKey string

	PublicKey string

	Config string
}

func NewProfile(
	email string,
	server string,
	ip string,
	privateKey string,
	publicKey string,
	config string,
) VPNProfile {

	return VPNProfile{

		Email: email,

		Server: server,

		IP: ip,

		PrivateKey: privateKey,

		PublicKey: publicKey,

		Config: config,
	}
}
