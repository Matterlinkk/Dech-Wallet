package main

import (
	"Signature/operations"
	"fmt"
	"math/big"
)

func main() {
	privateKey := 2

	x1 := "55066263022277343669578718895168534326250603453777594175500187360389116729240"
	x, _ := new(big.Int).SetString(x1, 10)

	y1 := "32670510020758816978083085130507043184471273380659243275938904335757337482424"
	y, _ := new(big.Int).SetString(y1, 10)

	gpPoint, err := operations.CreatePoint(x, y)

	publicKey := operations.Multiply(gpPoint, privateKey)

	fmt.Println("public: ", publicKey)
	fmt.Println("private: ", privateKey)

	signature, err := operations.SignMessage("Hello world", big.NewInt(int64(privateKey)))
	if err != nil {
		fmt.Errorf("Error %s: ", err)
	}
	fmt.Println(signature.R, signature.S)

	fmt.Println(operations.VerifySignature(signature, "Hello world", publicKey))

}
