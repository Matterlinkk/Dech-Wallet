package keys

import (
	"github.com/Matterlinkk/Dech-Wallet/point"
	"math/big"
)

type KeyPair struct {
	privateKey *big.Int
	publicKey  *point.Point
}

func GetKeys(privateKey *big.Int) KeyPair {

	gPoint := point.CreateGPoint()

	publicKey := gPoint.Multiply(privateKey)

	return KeyPair{
		privateKey: privateKey,
		publicKey:  publicKey,
	}
}

func (keys *KeyPair) GetPublic() *point.Point {
	return keys.publicKey
}

func (keys *KeyPair) GetPrivate() *big.Int {
	return keys.privateKey
}

func GetSharedSecret(publicKey *point.Point, privateKey *big.Int) *big.Int {

	sharedSecret := publicKey.Multiply(privateKey)

	return new(big.Int).Set(sharedSecret.GetX())
}
