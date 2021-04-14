package base

import (
	ecc "crypto/elliptic"
	"crypto/rand"
	"os"
)

func main() {
	k := generateRandom()
	x := generateRandom()
	_, X, Y, err := ecc.GenerateKey(ecc.P384(), rand.Reader)
	if err != nil {
		os.Exit(1)
	}
	K := ScalerMul(k, &Point{x: X, y: Y})

	hash := Hash(a, ScalerMul(a, &Point{x: x, y: y}))

}
