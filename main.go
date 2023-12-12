package main

import (
	"fmt"
	"github.com/Matterlinkk/Dech-Wallet/keys"
	"github.com/Matterlinkk/Dech-Wallet/signature"
	"math/big"
)

func main() {
	alice := keys.GetKeys(big.NewInt(1))
	fmt.Println(alice.PublicKey.GetAdress())

	message := "q"

	sign := signature.SignMessage(message, alice)

	fmt.Println(signature.VerifySignature(*sign, message, alice.PublicKey.PublicKey))
}
