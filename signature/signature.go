package signature

import (
	"fmt"
	"github.com/Matterlinkk/Dech-Wallet/hash"
	"github.com/Matterlinkk/Dech-Wallet/keys"
	"github.com/Matterlinkk/Dech-Wallet/operations"
	"github.com/Matterlinkk/Dech-Wallet/point"
	"math/big"
)

type Signature struct {
	owner *point.Point
	r     *big.Int
	s     *big.Int
}

func CreateSignature(part1, part2 *big.Int, publicKey *point.Point) *Signature {
	return &Signature{
		owner: publicKey,
		r:     part1,
		s:     part2,
	}
}

func (signature *Signature) GetSignature() {
	fmt.Printf("R: %s\nS: %s\n", signature.GetR().String(), signature.GetS().String())
}

func (signature *Signature) GetR() *big.Int {
	return signature.r
}

func (signature *Signature) GetS() *big.Int {
	return signature.s
}

func (signature *Signature) GetOwner() *point.Point {
	return signature.owner
}

func SignMessage(message string, keys keys.KeyPair) *Signature {

	k, _ := operations.GenerateRandomNumber()

	gPoint := point.CreateGPoint()

	kG := gPoint.Multiply(k)

	n1 := "115792089237316195423570985008687907852837564279074904382605163141518161494337" //115792089237316195423570985008687907852837564279074904382605163141518161494337 value from GP
	n, successN := new(big.Int).SetString(n1, 10)
	if !successN {
		panic("Error setting y value")
	}

	r := new(big.Int).Mod(kG.GetX(), n)

	if r.Cmp(big.NewInt(0)) == 0 {
		return SignMessage(message, keys)
	}

	hash := hash.SHA1(message)
	hashInt := new(big.Int).SetBytes(hash[:])

	// k^-1 * ( intHASH(message) + d * r) mod n, if s = 0 then do recursion
	invK := operations.FindInverse(k, n)
	dr := new(big.Int).Mul(keys.GetPrivate(), r)

	hashdr := new(big.Int).Add(hashInt, dr)

	kpandhash := new(big.Int).Mul(invK, hashdr)

	s := new(big.Int).Mod(kpandhash, n)
	if s.Cmp(big.NewInt(0)) == 0 {
		return SignMessage(message, keys)
	}

	return CreateSignature(r, s, keys.GetPublic())

}

func (signature *Signature) VerifySignature(message string, publicKey *point.Point) bool {

	n1 := "115792089237316195423570985008687907852837564279074904382605163141518161494337" //115792089237316195423570985008687907852837564279074904382605163141518161494337 value from GP
	n, successN := new(big.Int).SetString(n1, 10)
	if !successN {
		panic("Error setting y value")
	}

	sInverse := operations.FindInverse(signature.GetS(), n)
	sInverse.Mod(sInverse, n)

	// Calculate the hashed message
	hashedMessage := hash.SHA1(message)
	messageInt := new(big.Int).SetBytes(hashedMessage[:])

	// Calculate u and v
	u := new(big.Int).Mul(messageInt, sInverse)
	u.Mod(u, n)

	v := new(big.Int).Mul(signature.GetR(), sInverse)
	v.Mod(v, n)

	// Calculate the curve point P = u * G + v * publicKey
	gPoint := point.CreateGPoint()

	// Calculate u * G
	cPoint := gPoint.Multiply(u)

	// Calculate v * publicKey
	vPublicKey := publicKey.Multiply(v)

	// Calculate P = uG + vPublicKey
	p := cPoint.Add(vPublicKey)

	// Check if R is equal to x-coordinate of the point P
	return p.GetX().Cmp(signature.GetR()) == 0
}