package nodes

import (
	ecc "crypto/elliptic"
	"fmt"
	"math/big"

	base "github.com/ram-nad/crypto-project/base"
	chash "github.com/ram-nad/crypto-project/chameleon_hash"
)

func Test() {
	p := big.NewInt(11)
	n := big.NewInt(7)
	b := big.NewInt(7)

	curve := &ecc.CurveParams{Name: "Test"}
	curve.BitSize = 3
	curve.B = b
	curve.N = n
	curve.P = p
	curve.Gx = big.NewInt(1)
	curve.Gy = big.NewInt(4)

	bs := NewBaseStation(p, n, curve, base.NewPoint(curve.Gx, curve.Gy))
	nd := bs.AddNewNode()

	fmt.Printf("%v\n", bs)
	fmt.Printf("%v\n", nd)

	m := base.Concatenate(nd.N, nd.kP.X(), nd.n)
	fmt.Printf("%v\n", chash.GenerateHash(m, nd.r, nd.n, nd.kP, nd.Y, nd.Curve))
}
