package status

type VPNStatus struct {
	Email string

	Status string

	Server string

	IP string
}

func Connected(
	email string,
	server string,
	ip string,
) VPNStatus {

	return VPNStatus{

		Email: email,

		Status: "connected",

		Server: server,

		IP: ip,
	}
}

func Disconnected(
	email string,
) VPNStatus {

	return VPNStatus{

		Email: email,

		Status: "disconnected",
	}
}
