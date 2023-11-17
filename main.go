package main

import (
	"Signature/operations"
	"fmt"
	"math/big"
)

func main() {

	n1 := "115792089237316195423570985008687907852837564279074904382605163141518161494337" //n value from GP
	n, successN := new(big.Int).SetString(n1, 10)                                          //
	if !successN {
		panic("Error setting y value")
	}
	gpPoint, _ := operations.CreateGPoint()
	publicKey := operations.Multiply(gpPoint, n)

	fmt.Println(publicKey.X, publicKey.Y)
	//str := "1111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111010111010101011101101110011100110101011110100100010100000001110111011111111010010010111101000110011010000001101100100000101000001"
	//fmt.Println(string(str[255]), len(str))
}
