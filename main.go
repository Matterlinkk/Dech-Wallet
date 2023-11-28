package main

import (
	"Signature/operations"
	"fmt"
	"math/big"
)

func main() {
	message := "Hello, world!"

	alicePrivate := big.NewInt(123)
	bobPrivate := big.NewInt(321)

	aliceKeys := operations.GetKeyPair(alicePrivate)
	bobKeys := operations.GetKeyPair(bobPrivate)

	v1 := operations.GetSharedSecret(aliceKeys.PublicKey, bobKeys.PrivateKey)

	v2 := operations.GetSharedSecret(bobKeys.PublicKey, aliceKeys.PrivateKey)

	cipher := operations.GetEncryptedMessage(v1, message)

	decrypt := operations.GetDecryptedMessage(v2, cipher)

	fmt.Println(cipher, "\n")
	fmt.Println(decrypt, "\n")

}
