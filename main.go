package main

import (
	"Signature/operations"
	"fmt"
	"math/big"
)

func main() {

	privateKey := 1488
	message := "Hello, world!"
	gpPoint, _ := operations.CreateGPoint()
	publicKey := operations.Multiply(gpPoint, privateKey)

	signature, err := operations.SignMessage(message, big.NewInt(int64(privateKey)))
	if err != nil {
		panic("Panic in SignMessage")
	}

	fmt.Println(signature)

	fmt.Println(operations.VerifySignature(signature, message, publicKey))
}
