package main

import (
	"fmt"
	"math/big"
	"strconv"
)

func main() {
	answer := new(big.Int).Exp(big.NewInt(2), big.NewInt(1000), nil).String()

	acc := 0
	for i := 0; i < len(answer); i++ {
		foo, _ := strconv.Atoi(string(answer[i]))
		acc += foo
	}

	fmt.Println(answer)
	fmt.Println(acc)
}
