package main

import (
	"fmt"
	"github.com/Matterlinkk/Dech-Node/user"
	"github.com/Matterlinkk/Dech-Wallet/keys"
	"github.com/Matterlinkk/Dech-Wallet/operations"
	"math/big"
)

func main() {
	alice := user.CreateUser(big.NewInt(1))
	bob := user.CreateUser(big.NewInt(1))

	message := "Alice`s message"

	secretA := keys.GetSharedSecret(bob.PublicKey, alice.PrivateKey)

	secretB := keys.GetSharedSecret(alice.PublicKey, bob.PrivateKey)

	fmt.Printf("Alice`s secret: %s\nBob`s secret: %s\nSecrets are %t\n\n", secretA.String(), secretB.String(), secretA.String() == secretB.String())

	encryptedMessage := operations.GetEncryptedString(secretA.Bytes(), message)

	decryptedMessage := operations.GetDecryptedString(secretA.Bytes(), encryptedMessage)

	fmt.Printf("Original message: %s\nEncrypted message: %s\nDecrypted message: %s\n", message, encryptedMessage, decryptedMessage)

}
