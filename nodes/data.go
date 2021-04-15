package nodes

import (
	ecc "crypto/elliptic"
	"math/big"

	base "github.com/ram-nad/crypto-project/base"
	chash "github.com/ram-nad/crypto-project/chameleon_hash"
)

type NodeData struct {
	r     *big.Int
	k     *big.Int
	kP    *base.Point
	kXinv *big.Int
	N     *big.Int
	CH    *base.Point
}

type Nodes struct {
	r     *big.Int
	kP    *base.Point
	N     *big.Int
	CH    *base.Point
	a     *big.Int
	Curve ecc.Curve
	Y     *base.Point
	P     *base.Point
	q     *big.Int
	n     *big.Int
}

type BaseStation struct {
	x         *big.Int
	alpha     *big.Int
	Curve     ecc.Curve
	q         *big.Int
	n         *big.Int
	P         *base.Point
	Y         *base.Point
	CH        *base.Point
	nodesData []*NodeData
}

func NewBaseStation(q, n *big.Int, Curve ecc.Curve, P *base.Point) *BaseStation {
	x := base.GenerateRandom(n)
	alpha := base.GenerateRandom(n)
	nodesData := make([]*NodeData, 0)

	Y := base.ScalerMul(x, P, Curve)
	CH := base.ScalerMul(alpha, Y, Curve)

	return &BaseStation{x: x, alpha: alpha, Curve: Curve, q: q, n: n, P: P, Y: Y, CH: CH, nodesData: nodesData}
}

func NewNode(r, N, q, n *big.Int, CH, kP, Y, P *base.Point, Curve ecc.Curve) *Nodes {
	a := base.GenerateRandom(n)
	return &Nodes{r: r, kP: kP, a: a, q: q, n: n, N: N, CH: CH, Y: Y, P: P, Curve: Curve}
}

func (b *BaseStation) String() string {
	s := "BaseStation(x: " + b.x.String() + " alpha: " + b.alpha.String() + " Y: " + b.Y.String() + " CH: " + b.CH.String() + " P: " + b.P.String() + " q: " + b.q.String() + " n: " + b.n.String() + ")"
	for _, n := range b.nodesData {
		s += "\n\t\t" + n.String()
	}
	return s
}

func (n *NodeData) String() string {
	return "NodeData(N: " + n.N.String() + " k: " + n.k.String() + " r: " + n.r.String() + " CH: " + n.CH.String() + " kXinv: " + n.kXinv.String() + " kP: " + n.kP.String() + ")"
}

func (n *Nodes) String() string {
	return "Node(N: " + n.N.String() + " r: " + n.r.String() + " kP: " + n.kP.String() + " CH: " + n.CH.String() + " P: " + n.P.String() + " Y: " + n.Y.String() + " a: " + n.a.String() + " q: " + n.q.String() + " n: " + n.n.String() + ")"
}

func (b *BaseStation) AddNewNode() *Nodes {
	i := len(b.nodesData)
	N := big.NewInt(int64(i))
	k := base.GenerateRandom(b.n)
	kP := base.ScalerMul(k, b.P, b.Curve)
	m := base.Concatenate(N, kP.X(), b.n)
	kXinv := base.Div(k, b.x, b.q)
	r := chash.GenR(m, b.alpha, kXinv, b.q, b.n, kP)

	ND := &NodeData{r: r, k: k, kP: kP, kXinv: kXinv, N: N, CH: b.CH}

	b.nodesData = append(b.nodesData, ND)

	return NewNode(r, N, b.q, b.n, b.CH, kP, b.Y, b.P, b.Curve)
}
