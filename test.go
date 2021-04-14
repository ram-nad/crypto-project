package main

import (
	base "https://github.com/ram-nad/crypto-project/base"
	ecc "crypto/elliptic"
	"crypto/rand"
	"fmt"
	"math/big"
	"os"
)

func s() {
	m := big.NewInt(101)
	r := big.NewInt(11)
	k := base.generateRandom()
	x := base.generateRandom()
	_, ptX, ptY, err := ecc.GenerateKey(ecc.P384(), rand.Reader)
	if err != nil {
		os.Exit(1)
	}
	K := base.ScalerMul(k, &base.Point{x: ptX, y: ptY})
	Y := base.ScalerMul(x, &base.Point{x: ptX, y: ptY})
	pt_tmp := base.AddPoints(K, base.ScalerMul(r, Y))
	hash := base.ScalerMul(base.Hash(m, K), (pt_tmp))
	fmt.Println(hash)
}
