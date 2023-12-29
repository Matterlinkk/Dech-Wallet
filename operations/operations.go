package operations

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"github.com/tyler-smith/go-bip39"
	"io"
	"log"
	"math/big"
)

func CreateMnemonic() string {
	// Random entropy generation
	entropy, err := bip39.NewEntropy(128)
	if err != nil {
		log.Panicf("Error %s", err)
	}
	// Converting entropy into a mnemonic phrase
	mnemonic, err := bip39.NewMnemonic(entropy)
	if err != nil {
		log.Panicf("Error %s", err)
	}
	return mnemonic
}

func FindInverse(number, modulus *big.Int) *big.Int {
	// Calculate the modular inverse using exponentiation
	// The modular inverse of 'number' modulo 'modulus' is equivalent to 'number' raised to the power of 'modulus-2' modulo 'modulus'
	// This is based on Fermat's Little Theorem for prime 'modulus'
	exponent := new(big.Int).Sub(modulus, big.NewInt(2))
	inverse := new(big.Int).Exp(number, exponent, modulus)

	return inverse
}

func GenerateRandomNumber() (*big.Int, error) {

	n1 := "115792089237316195423570985008687907852837564279074904382605163141518161494337" //115792089237316195423570985008687907852837564279074904382605163141518161494337 value from GP
	n, successN := new(big.Int).SetString(n1, 10)
	n.Sub(n, big.NewInt(1))
	if !successN {
		panic("Error setting y value")
	}

	random, err := rand.Int(rand.Reader, n)
	if err != nil {
		return nil, err
	}

	return random, nil
}

func encrypt(plaintext []byte, block cipher.Block) []byte {

	plaintext = append(plaintext, bytes.Repeat([]byte{byte(16 - len(plaintext)%16)}, 16-len(plaintext)%16)...)
	ciphertext := make([]byte, len(plaintext))
	iv := make([]byte, aes.BlockSize)
	if _, err := io.ReadFull(rand.Reader, iv); err != nil {
		panic(err.Error())
	}

	mode := cipher.NewCBCEncrypter(block, iv)
	mode.CryptBlocks(ciphertext, plaintext)

	// add the IV to the beginning of the ciphertext
	ciphertext = append(iv, ciphertext...)

	return ciphertext
}

func GetEncryptedString(secret []byte, message string) string {

	// Convert a shared secret into a byte array of the required length for an AES key
	sharedSecretBytes := secret

	// AES random key generation
	key := make([]byte, 32)
	copy(key, sharedSecretBytes)

	// Initialization of AES block cipher
	block, err := aes.NewCipher(key)
	if err != nil {
		panic(err.Error())
	}

	// Message encryption
	plaintext := []byte(message)

	return string(encrypt(plaintext, block))

}

func decrypt(ciphertext []byte, block cipher.Block) []byte {
	if len(ciphertext) < aes.BlockSize {
		panic("ciphertext too short")
	}

	// Highlight the IV from the beginning of the ciphertext
	iv := ciphertext[:aes.BlockSize]
	ciphertext = ciphertext[aes.BlockSize:]

	mode := cipher.NewCBCDecrypter(block, iv)
	mode.CryptBlocks(ciphertext, ciphertext)

	// Remove the indentation added during encryption
	padSize := int(ciphertext[len(ciphertext)-1])
	return ciphertext[:len(ciphertext)-padSize]
}

func GetDecryptedString(secret []byte, ciphertext string) string {
	// Convert a shared secret into a byte array of the required length for an AES key
	sharedSecretBytes := secret

	// AES random key generation
	key := make([]byte, 32)
	copy(key, sharedSecretBytes)

	ciphertextByte := []byte(ciphertext)

	block, err := aes.NewCipher(key)
	if err != nil {
		panic(err.Error())
	}

	decryptedText := decrypt(ciphertextByte, block)

	return string(decryptedText)
}
