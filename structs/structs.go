package structs

import "math/big"

type Config struct {
	A big.Int
	B big.Int
	P big.Int
}

type Point struct {
	X *big.Int
	Y *big.Int
}

type Signature struct {
	R *big.Int
	S *big.Int
}
