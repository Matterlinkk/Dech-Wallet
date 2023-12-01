package operations

import (
	"Dech-Wallet/hash"
	"Dech-Wallet/structs"
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"fmt"
	"github.com/tyler-smith/go-bip39"
	"io"
	"log"
	"math/big"
)

func createPoint(x, y *big.Int) (*structs.Point, error) {
	config := DefaultConfig()

	ySquaredModP := new(big.Int).Mod(new(big.Int).Exp(y, big.NewInt(2), &config.P), &config.P)

	// (x^3 + a*x + b) % p
	rightSide := new(big.Int).Mod(new(big.Int).Add(
		new(big.Int).Exp(x, big.NewInt(3), &config.P),
		new(big.Int).Add(new(big.Int).Mul(&config.A, x), &config.B),
	), &config.P)

	if ySquaredModP.Cmp(rightSide) != 0 {
		return nil, fmt.Errorf("The point is not on the curve")
	}

	return &structs.Point{
		X: x,
		Y: y,
	}, nil
}

func DefaultConfig() *structs.Config {

	numberString := "115792089237316195423570985008687907853269984665640564039457584007908834671663"
	bigNumber := new(big.Int)
	bigNumber.SetString(numberString, 10)

	return &structs.Config{
		A: *big.NewInt(0),
		B: *big.NewInt(7),
		P: *bigNumber,
	}
}

func createSignature(part1, part2 *big.Int, publicKey *structs.Point) *structs.Signature {
	return &structs.Signature{
		Owner: publicKey,
		R:     part1,
		S:     part2,
	}
}

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

func isEqualTo(point1, point2 structs.Point) bool {
	return point1.X.Cmp(point2.X) == 0 && point1.Y.Cmp(point2.Y) == 0
}

func findInverse(number, modulus *big.Int) *big.Int {
	// Calculate the modular inverse using exponentiation
	// The modular inverse of 'number' modulo 'modulus' is equivalent to 'number' raised to the power of 'modulus-2' modulo 'modulus'
	// This is based on Fermat's Little Theorem for prime 'modulus'
	exponent := new(big.Int).Sub(modulus, big.NewInt(2))
	inverse := new(big.Int).Exp(number, exponent, modulus)

	return inverse
}

func doublePoint(point *structs.Point) *structs.Point {
	config := DefaultConfig()

	// s = (3 * x^2 + A) / (2 * y)
	numerator := new(big.Int).Mul(big.NewInt(3), new(big.Int).Exp(point.X, big.NewInt(2), &config.P))
	numerator.Add(numerator, &config.A)
	denominator := new(big.Int).Mul(big.NewInt(2), point.Y)
	inverse := findInverse(denominator, &config.P)
	slope := new(big.Int).Mul(numerator, inverse)
	slope.Mod(slope, &config.P)

	// x' = s^2 - 2 * x
	xPrime := new(big.Int).Exp(slope, big.NewInt(2), &config.P)
	xPrime.Sub(xPrime, new(big.Int).Mul(big.NewInt(2), point.X))
	xPrime.Mod(xPrime, &config.P)

	// y' = s * (x - x') - y
	yPrime := new(big.Int).Mul(slope, new(big.Int).Sub(point.X, xPrime))
	yPrime.Sub(yPrime, point.Y)
	yPrime.Mod(yPrime, &config.P)

	return &structs.Point{X: xPrime, Y: yPrime}
}

func add(point1, point2 *structs.Point) *structs.Point {
	config := DefaultConfig()

	if isEqualTo(*point1, *point2) {
		return doublePoint(point1)
	}

	deltaX := new(big.Int).Sub(point2.X, point1.X)
	deltaY := new(big.Int).Sub(point2.Y, point1.Y)
	inverse := findInverse(deltaX, &config.P)

	slope := new(big.Int).Mul(deltaY, inverse)
	slope.Mod(slope, &config.P)

	x := new(big.Int).Exp(slope, big.NewInt(2), &config.P)
	x.Sub(x, point2.X)
	x.Sub(x, point1.X)
	x.Mod(x, &config.P)

	y := new(big.Int).Mul(slope, new(big.Int).Sub(point1.X, x))
	y.Sub(y, point1.Y)
	y.Mod(y, &config.P)

	return &structs.Point{X: x, Y: y}
}

func multiply(point *structs.Point, times *big.Int) *structs.Point {
	result, _ := createPoint(point.X, point.Y)
	binTimes := fmt.Sprintf("%b", times)

	for i := 1; i < len(binTimes); i++ {
		result = doublePoint(result)

		if binTimes[i] == '1' {
			result = add(point, result)
		}
	}

	return result
}

func CreateGPoint() (*structs.Point, error) {
	x1 := "55066263022277343669578718895168534326250603453777594175500187360389116729240"
	x, successX := new(big.Int).SetString(x1, 10)
	if !successX {
		panic("Error setting x value")
	}

	y1 := "32670510020758816978083085130507043184471273380659243275938904335757337482424"
	y, successY := new(big.Int).SetString(y1, 10)
	if !successY {
		panic("Error setting y value")
	}

	return &structs.Point{
		X: x,
		Y: y,
	}, nil
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

func SignMessage(message string, keys structs.KeyPair) (*structs.Signature, error) {

	k, _ := GenerateRandomNumber()

	gPoint, err := CreateGPoint()
	if err != nil {
		log.Panicf("Error with GPoint: %s", err)
	}

	kG := multiply(gPoint, k)

	n1 := "115792089237316195423570985008687907852837564279074904382605163141518161494337" //115792089237316195423570985008687907852837564279074904382605163141518161494337 value from GP
	n, successN := new(big.Int).SetString(n1, 10)
	if !successN {
		panic("Error setting y value")
	}

	r := new(big.Int).Mod(kG.X, n)

	if r.Cmp(big.NewInt(0)) == 0 {
		return SignMessage(message, keys)
	}

	hash := hash.SHA1(message)
	hashInt := new(big.Int).SetBytes(hash[:])

	// k^-1 * ( intHASH(message) + d * r) mod n, if s = 0 then do recursion
	invK := findInverse(k, n)
	dr := new(big.Int).Mul(keys.PrivateKey, r)

	hashdr := new(big.Int).Add(hashInt, dr)

	kpandhash := new(big.Int).Mul(invK, hashdr)

	s := new(big.Int).Mod(kpandhash, n)
	if s.Cmp(big.NewInt(0)) == 0 {
		return SignMessage(message, keys)
	}

	return createSignature(r, s, keys.PublicKey), err

}

func VerifySignature(signature *structs.Signature, message string, publicKey *structs.Point) bool {

	n1 := "115792089237316195423570985008687907852837564279074904382605163141518161494337" //115792089237316195423570985008687907852837564279074904382605163141518161494337 value from GP
	n, successN := new(big.Int).SetString(n1, 10)
	if !successN {
		panic("Error setting y value")
	}

	sInverse := findInverse(signature.S, n)
	sInverse.Mod(sInverse, n)

	// Calculate the hashed message
	hashedMessage := hash.SHA1(message)
	messageInt := new(big.Int).SetBytes(hashedMessage[:])

	// Calculate u and v
	u := new(big.Int).Mul(messageInt, sInverse)
	u.Mod(u, n)

	v := new(big.Int).Mul(signature.R, sInverse)
	v.Mod(v, n)

	// Calculate the curve point P = u * G + v * publicKey
	gPoint, _ := CreateGPoint()

	// Calculate u * G
	cPoint := multiply(gPoint, u)

	// Calculate v * publicKey
	vPublicKey := multiply(publicKey, v)

	// Calculate P = uG + vPublicKey
	p := add(cPoint, vPublicKey)

	// Check if R is equal to x-coordinate of the point P
	return p.X.Cmp(signature.R) == 0
}

func GetKeyPair(privateKey *big.Int) structs.KeyPair {

	gPoint, _ := CreateGPoint()

	publicKey := multiply(gPoint, privateKey)

	return structs.KeyPair{
		PrivateKey: privateKey,
		PublicKey:  publicKey,
	}
}

func GetSharedSecret(publicKey *structs.Point, privateKey *big.Int) *big.Int {

	sharedSecret := multiply(publicKey, privateKey)

	return new(big.Int).Set(sharedSecret.X)
}

func encrypt(plaintext []byte, block cipher.Block) []byte {
	// Добавляем отступ до размера блока
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

func GetEncryptedMessage(secret *big.Int, message string) string {

	// Convert a shared secret into a byte array of the required length for an AES key
	sharedSecretBytes := secret.Bytes()

	// AES random key generation
	key := make([]byte, 32)
	copy(key, sharedSecretBytes) // Используем общий секрет в качестве ключа

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

func GetDecryptedMessage(secret *big.Int, ciphertext string) string {
	// Convert a shared secret into a byte array of the required length for an AES key
	sharedSecretBytes := secret.Bytes()

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
