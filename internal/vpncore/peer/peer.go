package peer

import (
	"crypto/rand"
	"encoding/base64"
)

type Peer struct {
	Email      string
	PublicKey  string
	PrivateKey string
	IP         string
}

func GenerateKeys() (string, string) {

	private := make([]byte, 32)

	rand.Read(private)

	privateKey := base64.StdEncoding.EncodeToString(
		private,
	)

	public := make([]byte, 32)

	rand.Read(public)

	publicKey := base64.StdEncoding.EncodeToString(
		public,
	)

	return privateKey, publicKey
}

func CreatePeer(
	email string,
	ip string,
) Peer {

	privateKey, publicKey := GenerateKeys()

	return Peer{

		Email: email,

		PrivateKey: privateKey,

		PublicKey: publicKey,

		IP: ip,
	}
}
