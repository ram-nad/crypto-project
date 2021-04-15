package chameleon_hash

import (
	ecc "crypto/elliptic"
	"math/big"

	base "github.com/ram-nad/crypto-project/base"
)

func GenerateHash(m, r, n *big.Int, kP, Y *base.Point, Curve ecc.Curve) *base.Point {
	f := base.Hash(m, kP, n)
	return base.AddPoints(base.ScalerMul(f, kP, Curve), base.ScalerMul(r, Y, Curve), Curve)
}

func GenR(m, alpha, kXinv, q, n *big.Int, kP *base.Point) *big.Int {
	f := base.Hash(m, kP, n)
	return base.Sub(alpha, base.Mul(f, kXinv, q), q)
}
