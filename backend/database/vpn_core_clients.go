package database

type VPNCoreClient struct {
	Email string

	IP string

	PublicKey string

	PrivateKey string

	Status string
}

func CreateVPNCoreClient(
	client interface{},
) error {

	return nil
}
