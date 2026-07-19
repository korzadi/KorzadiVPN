package wireguard

func CreateProfile(email string, server string) Client {

	client := NewClient(email)

	client.Server = server
	client.PrivateKey = GenerateKey()
	client.PublicKey = GenerateKey()

	return client
}
