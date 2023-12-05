package point

import (
	"fmt"
	"github.com/Matterlinkk/Dech-Wallet/config"
	"github.com/Matterlinkk/Dech-Wallet/operations"
	"math/big"
)

type Point struct {
	x *big.Int
	y *big.Int
}

func createPoint(x, y *big.Int) (*Point, error) {
	cfg := config.DefaultConfig()

	ySquaredModP := new(big.Int).Mod(new(big.Int).Exp(y, big.NewInt(2), cfg.GetP()), cfg.GetP())

	// (x^3 + a*x + b) % p
	rightSide := new(big.Int).Mod(new(big.Int).Add(
		new(big.Int).Exp(x, big.NewInt(3), cfg.GetP()),
		new(big.Int).Add(new(big.Int).Mul(cfg.GetA(), x), cfg.GetB()),
	), cfg.GetP())

	if ySquaredModP.Cmp(rightSide) != 0 {
		return nil, fmt.Errorf("The point is not on the curve")
	}

	return &Point{
		x: x,
		y: y,
	}, nil
}

func (p *Point) GetPoint() *Point {
	return p
}

func (p *Point) GetX() *big.Int {
	return p.x
}

func (p *Point) GetY() *big.Int {
	return p.y
}

func (p *Point) isEqualTo(point Point) bool {
	return p.GetX().Cmp(point.GetX()) == 0 && p.GetY().Cmp(point.GetY()) == 0
}

func CreateGPoint() *Point {
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

	return &Point{
		x: x,
		y: y,
	}
}

func (p *Point) DoublePoint() *Point {
	cfg := config.DefaultConfig()

	// s = (3 * x^2 + A) / (2 * y)
	numerator := new(big.Int).Mul(big.NewInt(3), new(big.Int).Exp(p.GetX(), big.NewInt(2), cfg.GetP()))
	numerator.Add(numerator, cfg.GetA())
	denominator := new(big.Int).Mul(big.NewInt(2), p.GetY())
	inverse := operations.FindInverse(denominator, cfg.GetP())
	slope := new(big.Int).Mul(numerator, inverse)
	slope.Mod(slope, cfg.GetP())

	// x' = s^2 - 2 * x
	xPrime := new(big.Int).Exp(slope, big.NewInt(2), cfg.GetP())
	xPrime.Sub(xPrime, new(big.Int).Mul(big.NewInt(2), p.GetX()))
	xPrime.Mod(xPrime, cfg.GetP())

	// y' = s * (x - x') - y
	yPrime := new(big.Int).Mul(slope, new(big.Int).Sub(p.GetX(), xPrime))
	yPrime.Sub(yPrime, p.GetY())
	yPrime.Mod(yPrime, cfg.GetP())

	return &Point{x: xPrime, y: yPrime}
}

func (p *Point) Add(point *Point) *Point {
	cfg := config.DefaultConfig()

	if p.isEqualTo(*point) {
		return p.DoublePoint()
	}

	deltaX := new(big.Int).Sub(point.GetX(), p.GetX())
	deltaY := new(big.Int).Sub(point.GetY(), p.GetY())
	inverse := operations.FindInverse(deltaX, cfg.GetP())

	slope := new(big.Int).Mul(deltaY, inverse)
	slope.Mod(slope, cfg.GetP())

	x := new(big.Int).Exp(slope, big.NewInt(2), cfg.GetP())
	x.Sub(x, point.GetX())
	x.Sub(x, p.GetX())
	x.Mod(x, cfg.GetP())

	y := new(big.Int).Mul(slope, new(big.Int).Sub(p.GetX(), x))
	y.Sub(y, p.GetY())
	y.Mod(y, cfg.GetP())

	return &Point{x: x, y: y}
}

func (p *Point) Multiply(times *big.Int) *Point {
	result, _ := createPoint(p.GetX(), p.GetY())
	binTimes := fmt.Sprintf("%b", times)

	for i := 1; i < len(binTimes); i++ {
		result = result.DoublePoint()

		if binTimes[i] == '1' {
			result = p.Add(result)
		}
	}

	return result
}
