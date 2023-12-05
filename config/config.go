package config

import (
	"math/big"
)

type Config struct {
	a big.Int
	b big.Int
	p big.Int
}

func DefaultConfig() *Config {

	numberString := "115792089237316195423570985008687907853269984665640564039457584007908834671663"
	bigNumber := new(big.Int)
	bigNumber.SetString(numberString, 10)

	return &Config{
		a: *big.NewInt(0),
		b: *big.NewInt(7),
		p: *bigNumber,
	}
}

func (c Config) GetA() *big.Int {
	return &c.a
}

func (c Config) GetB() *big.Int {
	return &c.b
}

func (c Config) GetP() *big.Int {
	return &c.p
}
