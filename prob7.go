package main

import (
	"fmt"
	"math"
)

func main() {
	primeCount := 1
	lastPrime := 0
	isPrime := true
	const primeToFind = 10001

	for i := 3; primeCount < primeToFind; i += 2 {
		isPrime = true
		upto := int(math.Sqrt(float64(i)))
		for j := 2; j <= upto; j++ {
			if i%j == 0 {
				isPrime = false
			}
		}
		if isPrime {
			primeCount++
			lastPrime = i
		}
	}
	fmt.Println(fmt.Sprintf("%dth prime is %d", primeCount, lastPrime))
}
