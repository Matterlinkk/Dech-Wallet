package testing

import (
	"Signature/operations"
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
