package main

import (
	"fmt"
	"github.com/Matterlinkk/Dech-Wallet/keys"
	"github.com/Matterlinkk/Dech-Wallet/signature"
	"math/big"
)

func main() {
	message := "qwe"

	moonkushPrivate := big.NewInt(1488)

	aliceKeys := keys.GetKeys(moonkushPrivate)

	sign := signature.SignMessage(message, aliceKeys)

	sign.GetSignature()

	q := signature.VerifySignature(*sign, "qwe", aliceKeys.GetPublic())
	fmt.Println(q)
}
