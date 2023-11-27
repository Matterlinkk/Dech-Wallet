package main

import (
	"Signature/operations"
	"fmt"
	"math/big"
)

func main() {

	alicePrivate := big.NewInt(123)
	bobPrivate := big.NewInt(321)

	aliceKeys := operations.GetKeyPair(alicePrivate)
	bobKeys := operations.GetKeyPair(bobPrivate)

	// Print intermediate values
	fmt.Println("Alice Private Key:", aliceKeys.PrivateKey)
	fmt.Println("Bob Private Key:", bobKeys.PrivateKey)
	fmt.Println("Alice Public Key:", aliceKeys.PublicKey)
	fmt.Println("Bob Public Key:", bobKeys.PublicKey)

	// Alice's calculation
	v1 := operations.GetSharedSecret(aliceKeys.PublicKey, bobKeys.PrivateKey)

	// Bob's calculation
	v2 := operations.GetSharedSecret(bobKeys.PublicKey, aliceKeys.PrivateKey)

	fmt.Println("Shared Secret Alice:", v1)
	fmt.Println("Shared Secret Bob:", v2)

	// Check if the shared secrets are equal
	if v1.Cmp(v2) == 0 {
		fmt.Println("Shared secrets match. Secure communication established.")
	} else {
		fmt.Println("Shared secrets do not match. Error in key exchange.")
	}
}

// gp x: 55066263022277343669578718895168534326250603453777594175500187360389116729240
// gp y: 32670510020758816978083085130507043184471273380659243275938904335757337482424
// p: 115792089237316195423570985008687907853269984665640564039457584007908834671663
// a: 0, b: 7
