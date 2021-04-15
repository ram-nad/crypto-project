package base

import (
	ecc "crypto/elliptic"
	hash "crypto/sha512"
	"math/big"
)

type Point struct {
	x, y *big.Int
}

func (p *Point) String() string {
	return "(" + p.x.String() + ", " + p.y.String() + ")"
}

func NewPoint(x, y *big.Int) *Point {
	return &Point{x: x, y: y}
}

func (p *Point) X() *big.Int {
	return p.x
}

func (p *Point) Y() *big.Int {
	return p.x
}

func IsIntEq(a, b *big.Int) bool {
	return a.Cmp(b) == 0
}

func IsPointEqual(a, b *Point) bool {
	return IsIntEq(a.x, b.x) && IsIntEq(a.y, b.y)
}

func PadBigInt(i *big.Int, size int) []byte {
	var a []byte = make([]byte, size)
	return i.FillBytes(a)
}

func Concatenate(a, b, N *big.Int) *big.Int {
	size := (N.BitLen() + 7) / 8
	x := PadBigInt(a, size)
	y := PadBigInt(b, size)

	c := append(x, y...)

	d := new(big.Int).SetBytes(c)

	return new(big.Int).Mod(d, N)
}

func Hash(a *big.Int, p *Point, N *big.Int) *big.Int {
	c := Concatenate(a, Concatenate(p.x, p.y, N), N)
	h := hash.Sum384(c.Bytes())

	d := new(big.Int).SetBytes(h[:])

	return new(big.Int).Mod(d, N)
}

func Add(a, b, PRIME *big.Int) *big.Int {
	c := new(big.Int)
	c.Add(a, b)
	return new(big.Int).Mod(c, PRIME)
}

func Sub(a, b, PRIME *big.Int) *big.Int {
	c := new(big.Int)
	c.Sub(a, b)
	return new(big.Int).Mod(c, PRIME)
}

func Mul(a, b, PRIME *big.Int) *big.Int {
	c := new(big.Int)
	c.Mul(a, b)
	return new(big.Int).Mod(c, PRIME)
}

func Inv(a, PRIME *big.Int) *big.Int {
	c := new(big.Int)
	return c.ModInverse(a, PRIME)
}

func Div(a, b, PRIME *big.Int) *big.Int {
	return Mul(a, Inv(b, PRIME), PRIME)
}

func AddPoints(a, b *Point, Curve ecc.Curve) *Point {
	x, y := Curve.Add(a.x, a.y, b.x, b.y)
	return &Point{x: x, y: y}
}

func ScalerMul(a *big.Int, p *Point, Curve ecc.Curve) *Point {
	x, y := Curve.ScalarMult(p.x, p.y, a.Bytes())
	return &Point{x: x, y: y}
}
