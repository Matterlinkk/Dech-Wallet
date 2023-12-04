package testing

import (
	"github.com/Matterlinkk/Dech-Wallet/operations"
	"math/big"
	"testing"
)

func TestSignature(t *testing.T) {

	alice := operations.GetKeyPair(big.NewInt(1))

	m := "Hello, ECDSA!"

	signature, _ := operations.SignMessage(m, alice)

	verify := operations.VerifySignature(signature, m, alice.PublicKey)

	if !verify {
		t.Error("Invalid signature")
	}
}
