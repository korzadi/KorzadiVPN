package allocator

type Assignment struct {
	Email string

	IP string

	Server string
}

func Assign(
	email string,
	server string,
	ip string,
) Assignment {

	return Assignment{

		Email: email,

		IP: ip,

		Server: server,
	}
}
