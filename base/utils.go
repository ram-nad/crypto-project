package base

import (
	hash "crypto/sha512"
	"math/big"
)

func PadBigInt(i *big.Int) []byte {
	var a []byte = make([]byte, SIZE)
	return i.FillBytes(a)
}

func Concatenate(a, b *big.Int) *big.Int {
	x := PadBigInt(a)
	y := PadBigInt(b)

	c := append(x, y...)

	d := new(big.Int).SetBytes(c)

	return new(big.Int).Mod(d, N)
}

func Hash(a *big.Int, p *Point) *big.Int {
	c := Concatenate(a, Concatenate(p.x, p.y))
	h := hash.Sum384(c.Bytes())

	d := new(big.Int).SetBytes(h[:])

	return new(big.Int).Mod(d, N)
}

func Add(a, b *big.Int) *big.Int {
	c := new(big.Int)
	c.Add(a, b)
	return new(big.Int).Mod(c, PRIME)
}

func Sub(a, b *big.Int) *big.Int {
	c := new(big.Int)
	c.Sub(a, b)
	return new(big.Int).Mod(c, PRIME)
}

func Mul(a, b *big.Int) *big.Int {
	c := new(big.Int)
	c.Mul(a, b)
	return new(big.Int).Mod(c, PRIME)
}

func Inv(a *big.Int) *big.Int {
	c := new(big.Int)
	return c.ModInverse(a, PRIME)
}

func Div(a, b *big.Int) *big.Int {
	return Mul(a, Inv(b))
}

func AddPoints(a, b *Point) *Point {
	x, y := CURVE.Add(a.x, a.y, a.x, a.y)
	return &Point{x: x, y: y}
}

func ScalerMul(a *big.Int, p *Point) *Point {
	x, y := CURVE.ScalarMult(p.x, p.y, a.Bytes())
	return &Point{x: x, y: y}
}
