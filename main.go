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

	//bobPrivate := big.NewInt(122)

	//bobKeys := operations.GetKeyPair(bobPrivate)

	fmt.Println(signature)
	fmt.Printf("R: %s\nS: %s\n", signature.R.String(), signature.S.String())

	q := operations.VerifySignature(signature, "Hello, world", aliceKeys.PublicKey)
	fmt.Println(q)
}
