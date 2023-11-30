package main

import (
	"Signature/operations"
	"fmt"
	"math/big"
)

func main() {
	message := "Hello, world!"

	alicePrivate := big.NewInt(123)

	aliceKeys := operations.GetKeyPair(alicePrivate)

	signature, _ := operations.SignMessage(message, aliceKeys)

	q := operations.VerifySignature(signature, message, aliceKeys.PublicKey)
	fmt.Println(q)
}
