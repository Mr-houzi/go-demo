package main

import (
	"fmt"
	"math"
	"math/big"
)

func main() {
	// int
	var ia int64 = math.MaxInt64
	var ib int64 = math.MaxInt64
	var iout int64 = ia * ib
	fmt.Printf("ia int64:%v\n", ia)
	fmt.Printf("ib int64:%v\n", ib)
	fmt.Printf("iout int64:%v\n", iout) // 溢出，计算结果不准

	// big.int
	bia := big.NewInt(math.MaxInt64)
	bib := big.NewInt(math.MaxInt64)
	biout := big.NewInt(1)
	biout.Mul(bia, bib).Mul(biout, bib)
	fmt.Printf("a Big Int: %v\n", bia)
	fmt.Printf("b Big Int: %v\n", bib)
	fmt.Printf("out Big Int: %v\n", biout.String()) // 结果超出 int64 上线，但依然能表示出来
}
