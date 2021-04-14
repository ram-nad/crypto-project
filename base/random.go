package base

import (
	"crypto/rand"
	"fmt"
	"math/big"
)

func generateRandom() *big.Int {
	b := make([]byte, SIZE)
	_, err := rand.Read(b)
	if err != nil {
		fmt.Println("error:", err)
		return nil
	}
	x := new(big.Int).SetBytes(b)
	return x.Mod(x, N)
}
