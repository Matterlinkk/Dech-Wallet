package keys

import (
	"github.com/Matterlinkk/Dech-Wallet/point"
	"github.com/Matterlinkk/Dech-Wallet/publickey"
	"math/big"
)

type KeyPair struct {
	PrivateKey *big.Int             `json:"privateKey"`
	PublicKey  *publickey.PublicKey `json:"publicKey"`
}

func GetKeys(privateKey *big.Int) KeyPair {

	gPoint := point.CreateGPoint()

	publicKey := gPoint.Multiply(privateKey)

	return KeyPair{
		PrivateKey: privateKey,
		PublicKey: &publickey.PublicKey{
			PublicKey: publicKey,
		},
	}
}

func GetSharedSecret(publicKey *point.Point, privateKey *big.Int) *big.Int {

	sharedSecret := publicKey.Multiply(privateKey)

	return new(big.Int).Set(sharedSecret.GetX())
}
