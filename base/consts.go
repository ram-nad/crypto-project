package base

import (
	ecc "crypto/elliptic"
	"math/big"
)

type Point struct {
	x, y *big.Int
}

var SIZE int = ecc.P384().Params().BitSize / 8

var PRIME *big.Int = ecc.P384().Params().P

var N *big.Int = ecc.P384().Params().N

var CURVE ecc.Curve = ecc.P384()

var BasePoint *Point = &Point{x: ecc.P384().Params().Gx, y: ecc.P384().Params().Gy}
