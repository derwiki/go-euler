package main

import (
	"fmt"
	"math"
)

func main() {
	primeCount := 0
	lastPrime := 0
	const primeToFind = 10001

	for i := 2; primeCount < primeToFind; i++ {
		upto := int(math.Sqrt(float64(i)))
		isPrime := true
		for j := 2; j <= upto; j++ {
			if i%j == 0 {
				isPrime = false
			}
		}
		if isPrime {
			primeCount++
			lastPrime = i
			fmt.Println(fmt.Sprintf("%dth prime is %d", primeCount, lastPrime))
		}
	}
}
