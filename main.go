package main

import (
	"fmt"
	"github.com/Matterlinkk/Dech-Wallet/keys"
	"math/big"
)

func main() {
	alice := keys.GetKeys(big.NewInt(1))
	fmt.Println(alice.PublicKey.GetAdress())
}
