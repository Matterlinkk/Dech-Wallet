package testing

import (
	"github.com/Matterlinkk/Dech-Wallet/keys"
	"github.com/Matterlinkk/Dech-Wallet/signature"
	"math/big"
	"testing"
)

func TestSignature(t *testing.T) {

	alice := keys.GetKeys(big.NewInt(1))

	m := "Hello, ECDSA!"

	sign := signature.SignMessage(m, alice)

	verify := sign.VerifySignature(m, alice.GetPublic())

	if !verify {
		t.Error("Invalid signature")
	}
}
