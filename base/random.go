package base

import (
	"crypto/rand"
	"fmt"
	"math/big"
)

func GenerateRandom(N *big.Int) *big.Int {
	size := (N.BitLen() + 7) / 8
	b := make([]byte, size)
	_, err := rand.Read(b)
	if err != nil {
		fmt.Println("error:", err)
		return nil
	}
	x := new(big.Int).SetBytes(b)
	y := new(big.Int).Mod(x, N)
	if y.Cmp(big.NewInt(0)) == 0 {
		return big.NewInt(1)
	}
	return y
}
