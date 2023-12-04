package main

import (
	"fmt"
	"github.com/Matterlinkk/Dech-Wallet/operations"
	"math/big"
)

func main() {
	message := "qwe"

	moonkushPrivate := big.NewInt(1488)

	aliceKeys := operations.GetKeyPair(moonkushPrivate)

	signature, _ := operations.SignMessage(message, aliceKeys)

	//bobPrivate := big.NewInt(122)

	//bobKeys := operations.GetKeyPair(bobPrivate)

	fmt.Println(signature)
	fmt.Printf("R: %s\nS: %s\n", signature.R.String(), signature.S.String())

	q := operations.VerifySignature(signature, "qwe", aliceKeys.PublicKey)
	fmt.Println(q)
}
