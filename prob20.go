package main

import (
	"fmt"
	"math/big"
	"strconv"
)

func main() {
	n := 100
	factorial := new(big.Int).MulRange(1, int64(n)).String()
	fmt.Println(factorial)
	acc := 0
	for i := 0; i < len(factorial); i++ {
		foo, _ := strconv.Atoi(string(factorial[i]))
		acc += foo
	}
	fmt.Println(acc)
}
