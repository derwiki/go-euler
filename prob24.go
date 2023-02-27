package main

// A permutation is an ordered arrangement of objects. For example, 3124 is one possible permutation of the digits 1, 2, 3 and 4. If all of the permutations are listed numerically or alphabetically, we call it lexicographic order. The lexicographic permutations of 0, 1 and 2 are:

// 012   021   102   120   201   210

// What is the millionth lexicographic permutation of the digits 0, 1, 2, 3, 4, 5, 6, 7, 8 and 9?

import (
	"fmt"
	"math/big"
	"sort"
	"strconv"
	"strings"
)

func main() {
	digits := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	sort.Ints(digits)
	n := len(digits)
	permutations := []int{}

	for i := 0; i < factorial(n); i++ {
		p := big.NewInt(int64(i))
		q := big.NewInt(int64(n))
		r := big.NewInt(0)
		perm := make([]int, n)

		for j := 0; j < n; j++ {
			r.Mod(p, q)
			k := int(r.Int64())
			p.Div(p, q)
			perm[j] = digits[k]
			digits = append(digits[:k], digits[k+1:]...)
			q.Sub(q, big.NewInt(1))
		}
		permAsInt, _ := strconv.Atoi(strings.Trim(strings.Join(strings.Fields(fmt.Sprint(perm)), ""), "[]"))
		permutations = append(permutations, permAsInt)

		digits = []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	}
	sort.Ints(permutations)
	fmt.Println(permutations[1_000_000-1])
}

func factorial(n int) int {
	if n == 0 {
		return 1
	}
	return n * factorial(n-1)
}
