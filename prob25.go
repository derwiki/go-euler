package main

import (
	"fmt"
	"math/big"
)

func main() {
	a := big.NewInt(int64(0))
	b := big.NewInt(int64(1))

	i := 1
	for ; len(b.String()) < 1000; i++ {
		//fmt.Println(i, a, b)
		sum := big.NewInt(0)
		sum.Add(a, b)
		a = b
		b = sum
	}
	fmt.Println(i, b)
}
