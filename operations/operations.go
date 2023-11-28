package operations

import (
	"Signature/hash"
	"Signature/structs"
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

func CreatePoint(x, y *big.Int) (*structs.Point, error) {
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

func CreateSignature(part1, part2 *big.Int) *structs.Signature {
	return &structs.Signature{
		R: part1,
		S: part2,
	}
}

func CreateMnemonic() string { //Dech-Wallet закинуть туда
	// Генерация случайной энтропии
	entropy, err := bip39.NewEntropy(128)
	if err != nil {
		log.Panicf("Error %s", err)
	}
	// Преобразование энтропии в мнемоническую фразу
	mnemonic, err := bip39.NewMnemonic(entropy)
	if err != nil {
		log.Panicf("Error %s", err)
	}
	return mnemonic
}

func IsEqualTo(point1, point2 structs.Point) bool {
	return point1.X.Cmp(point2.X) == 0 && point1.Y.Cmp(point2.Y) == 0
}

func FindInverse(number, modulus *big.Int) *big.Int {
	inverse := new(big.Int).ModInverse(number, modulus)

	return inverse
}

func DoublePoint(point *structs.Point) *structs.Point {
	config := DefaultConfig()

	// s = (3 * x^2 + A) / (2 * y)
	numerator := new(big.Int).Mul(big.NewInt(3), new(big.Int).Exp(point.X, big.NewInt(2), &config.P))
	numerator.Add(numerator, &config.A)
	denominator := new(big.Int).Mul(big.NewInt(2), point.Y)
	inverse := FindInverse(denominator, &config.P)
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

func Add(point1, point2 *structs.Point) *structs.Point {
	config := DefaultConfig()

	if IsEqualTo(*point1, *point2) {
		return DoublePoint(point1)
	}

	deltaX := new(big.Int).Sub(point2.X, point1.X)
	deltaY := new(big.Int).Sub(point2.Y, point1.Y)
	inverse := FindInverse(deltaX, &config.P)

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

func Multiply(point *structs.Point, times *big.Int) *structs.Point {
	result, _ := CreatePoint(point.X, point.Y)
	binTimes := fmt.Sprintf("%b", times)

	for i := 1; i < len(binTimes); i++ {
		result = DoublePoint(result)

		if binTimes[i] == '1' {
			result = Add(point, result)
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

func SignMessage(message string, privateKey *big.Int) (*structs.Signature, error) {

	k, _ := GenerateRandomNumber()
	gpPoint, err := CreateGPoint()
	if err != nil {
		log.Panicf("Error with GPoint: %s", err)
	}

	n1 := "115792089237316195423570985008687907852837564279074904382605163141518161494337" //115792089237316195423570985008687907852837564279074904382605163141518161494337 value from GP
	n, successN := new(big.Int).SetString(n1, 10)
	n.Sub(n, big.NewInt(1))
	if !successN {
		panic("Error setting y value")
	}
	fmt.Println("n is: ", n.String())

	rPoint := Multiply(gpPoint, n)

	fmt.Println("rPoint.X is: ", rPoint.X)

	r := new(big.Int).Mod(rPoint.X, n)

	fmt.Println("r is: ", r.String())

	if r.Cmp(big.NewInt(0)) == 0 {
		return SignMessage(message, privateKey)
	}

	// Create a copy of k to avoid modifying the original value
	kCopy := new(big.Int).Set(k)

	fmt.Println("kCopy: ", kCopy)
	fmt.Println("n: ", n)

	kInverse := FindInverse(kCopy, n)

	hashedMessage := hash.SHA1(message)
	messageInt := new(big.Int).SetBytes(hashedMessage[:])

	// Create a copy of privateKey to avoid modifying the original value
	privateKeyCopy := new(big.Int).Set(privateKey)
	fmt.Println("kInverse: ", kInverse)
	fmt.Println("messageInt: ", messageInt)
	fmt.Println("r: ", r)
	fmt.Println("privateKeyCopy: ", privateKeyCopy)
	s := new(big.Int).Mul(kInverse, new(big.Int).Add(messageInt, new(big.Int).Mul(r, privateKeyCopy)))
	println("s value: ", s.String())
	s.Mod(s, n)
	println("s value: ", s.String())

	return CreateSignature(r, s), nil
}

func VerifySignature(signature *structs.Signature, message string, publicKey *structs.Point) bool {
	r := signature.R
	s := signature.S

	n1 := "115792089237316195423570985008687907852837564279074904382605163141518161494337" //115792089237316195423570985008687907852837564279074904382605163141518161494337 value from GP
	n, successN := new(big.Int).SetString(n1, 10)
	n.Sub(n, big.NewInt(1))
	if !successN {
		panic("Error setting y value")
	}

	config := DefaultConfig()

	gpPoint, _ := CreateGPoint()

	sInverse := FindInverse(s, n)
	u := new(big.Int).SetBytes([]byte(message))
	u.Mul(u, sInverse)
	u.Mod(u, &config.P)

	v := new(big.Int).Set(r)
	v.Mul(v, sInverse)
	v.Mod(v, &config.P)

	cPoint := Add(Multiply(gpPoint, u), Multiply(publicKey, v)) //cPoint := Multiply(Multiply(gpPoint, int(u.Int64())), int(v.Int64())) 1.

	return cPoint.X.Cmp(r) == 0
}

func GetKeyPair(privateKey *big.Int) structs.KeyPair {

	gPoint, _ := CreateGPoint()

	publicKey := Multiply(gPoint, privateKey)

	return structs.KeyPair{
		PrivateKey: privateKey,
		PublicKey:  publicKey,
	}
}

func GetSharedSecret(publicKey *structs.Point, privateKey *big.Int) *big.Int {

	sharedSecret := Multiply(publicKey, privateKey)

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

	// Добавляем IV в начало зашифрованного текста
	ciphertext = append(iv, ciphertext...)

	return ciphertext
}

func GetEncryptedMessage(secret *big.Int, message string) string {

	// Преобразование общего секрета в массив байт нужной длины для ключа AES
	sharedSecretBytes := secret.Bytes()

	// Генерация случайного ключа AES
	key := make([]byte, 32)
	copy(key, sharedSecretBytes) // Используем общий секрет в качестве ключа

	// Инициализация блочного шифра AES
	block, err := aes.NewCipher(key)
	if err != nil {
		panic(err.Error())
	}

	// Шифрование сообщения
	plaintext := []byte(message)

	return string(encrypt(plaintext, block))

}

func decrypt(ciphertext []byte, block cipher.Block) []byte {
	if len(ciphertext) < aes.BlockSize {
		panic("ciphertext too short")
	}

	// Выделяем IV из начала зашифрованного текста
	iv := ciphertext[:aes.BlockSize]
	ciphertext = ciphertext[aes.BlockSize:]

	mode := cipher.NewCBCDecrypter(block, iv)
	mode.CryptBlocks(ciphertext, ciphertext)

	// Удаляем отступ, добавленный при шифровании
	padSize := int(ciphertext[len(ciphertext)-1])
	return ciphertext[:len(ciphertext)-padSize]
}

func GetDecryptedMessage(secret *big.Int, ciphertext string) string {
	// Преобразование общего секрета в массив байт нужной длины для ключа AES
	sharedSecretBytes := secret.Bytes()

	// Генерация случайного ключа AES
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
