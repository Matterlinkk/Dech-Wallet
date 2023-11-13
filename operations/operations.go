package operations

import (
	"Signature/hash"
	"Signature/structs"
	"crypto/rand"
	"fmt"
	"github.com/tyler-smith/go-bip39"
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
		X:      x,
		Y:      y,
		Config: config,
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
	// Используем функцию ModInverse для вычисления обратного элемента в модульном арифметическом кольце.
	// Если обратный элемент не существует, результат будет nil.
	inverse := new(big.Int).ModInverse(number, modulus)

	return inverse
}

func Add(point1, point2 *structs.Point) *structs.Point { //valid
	p := &point1.Config.P

	// Проверка на совпадение точек
	var slope *big.Int
	if IsEqualTo(*point1, *point2) {
		// Удвоение точки
		squareX := new(big.Int).Exp(point1.X, big.NewInt(2), p)
		numerator := new(big.Int).Mul(big.NewInt(3), squareX)
		denominator := new(big.Int).Mul(big.NewInt(2), point1.Y)
		inverse := FindInverse(denominator, p)

		slope = new(big.Int).Mul(numerator, inverse)
		slope.Mod(slope, p)
	} else {
		// Общий случай
		deltaX := new(big.Int).Sub(point2.X, point1.X)
		deltaY := new(big.Int).Sub(point2.Y, point1.Y)
		inverse := FindInverse(deltaX, p)

		slope = new(big.Int).Mul(deltaY, inverse)
		slope.Mod(slope, p)
	}

	// Вычисление новых координат
	x := new(big.Int).Exp(slope, big.NewInt(2), p)
	x.Sub(x, point2.X)
	x.Sub(x, point1.X)
	x.Mod(x, p)

	y := new(big.Int).Mul(slope, new(big.Int).Sub(point1.X, x))
	y.Sub(y, point1.Y)
	y.Mod(y, p)

	return &structs.Point{Config: point1.Config, X: x, Y: y}
}

func Multiply(point *structs.Point, times int) *structs.Point {
	currentPoint := point
	currentCoefficient := big.NewInt(1)

	previousPoints := make([]struct {
		coefficient *big.Int
		point       *structs.Point
	}, 0)

	for currentCoefficient.Cmp(big.NewInt(int64(times))) < 0 {
		// Save the current point in the list of previous points
		previousPoints = append(previousPoints, struct {
			coefficient *big.Int
			point       *structs.Point
		}{currentCoefficient, currentPoint})

		// If we can multiply the current point by 2, do so
		if new(big.Int).Mul(big.NewInt(2), currentCoefficient).Cmp(big.NewInt(int64(times))) <= 0 {
			currentPoint = Add(currentPoint, currentPoint)
			currentCoefficient.Mul(big.NewInt(2), currentCoefficient)
		} else {
			// Find the largest suitable point and add the current point to it
			var nextPoint *structs.Point = point
			nextCoefficient := big.NewInt(1)

			for _, previous := range previousPoints {
				if new(big.Int).Add(previous.coefficient, currentCoefficient).Cmp(big.NewInt(int64(times))) <= 0 {
					if previous.point.X.Cmp(currentPoint.X) != 0 {
						nextCoefficient = previous.coefficient
						nextPoint = previous.point
					}
				}
			}

			currentPoint = Add(currentPoint, nextPoint)
			currentCoefficient.Add(currentCoefficient, nextCoefficient)
		}
	}

	return currentPoint
}

func SignMessage(message string, privateKey *big.Int) (*structs.Signature, error) {
	x1 := "55066263022277343669578718895168534326250603453777594175500187360389116729240"
	x, successX := new(big.Int).SetString(x1, 10)
	if !successX {
		return nil, fmt.Errorf("Error setting x value")
	}

	y1 := "32670510020758816978083085130507043184471273380659243275938904335757337482424"
	y, successY := new(big.Int).SetString(y1, 10)
	if !successY {
		return nil, fmt.Errorf("Error setting y value")
	}

	gpPoint, err := CreatePoint(x, y)
	if err != nil {
		return nil, fmt.Errorf("Error creating point: %s", err)
	}

	k, err := rand.Int(rand.Reader, &gpPoint.Config.P)
	if err != nil {
		return nil, fmt.Errorf("Error generating random number k: %s", err)
	}

	rPoint := Multiply(gpPoint, int(k.Int64()))
	r := new(big.Int).Mod(rPoint.X, &gpPoint.Config.P)

	if r.Cmp(big.NewInt(0)) == 0 {
		// If r is zero, repeat the signing process
		return SignMessage(message, privateKey)
	}

	// Create a copy of k to avoid modifying the original value
	kCopy := new(big.Int).Set(k)

	kInverse := FindInverse(kCopy, &gpPoint.Config.P)

	hashedMessage := hash.SHA1(message)
	messageInt := new(big.Int).SetBytes(hashedMessage[:])

	// Create a copy of privateKey to avoid modifying the original value
	privateKeyCopy := new(big.Int).Set(privateKey)

	s := new(big.Int).Mul(kInverse, new(big.Int).Add(messageInt, new(big.Int).Mul(r, privateKeyCopy)))
	s.Mod(s, &gpPoint.Config.P)

	if s.Cmp(big.NewInt(0)) == 0 {
		// If s is zero, repeat the signing process
		return SignMessage(message, privateKey)
	}

	return CreateSignature(r, s), nil
}

func VerifySignature(signature *structs.Signature, message string, publicKey *structs.Point) bool {
	r := signature.R
	s := signature.S

	x1 := "55066263022277343669578718895168534326250603453777594175500187360389116729240"
	x, _ := new(big.Int).SetString(x1, 10)

	y1 := "32670510020758816978083085130507043184471273380659243275938904335757337482424"
	y, _ := new(big.Int).SetString(y1, 10)

	gpPoint, _ := CreatePoint(x, y)

	sInverse := FindInverse(s, &publicKey.Config.P)
	u := new(big.Int).SetBytes([]byte(message))
	u.Mul(u, sInverse)
	u.Mod(u, &publicKey.Config.P)

	v := new(big.Int).Set(r)
	v.Mul(v, sInverse)
	v.Mod(v, &publicKey.Config.P)

	cPoint := Multiply(Multiply(gpPoint, int(u.Int64())), int(v.Int64()))

	return cPoint.X.Cmp(r) == 0
}
