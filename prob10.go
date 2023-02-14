package main

import "fmt"

func sieveOfEratosthenes(limit int) []int {
	// Create a slice to store the prime numbers
	var primes []int
	// Create a slice to store the composites (non-prime numbers)
	composites := make([]bool, limit+1)
	// Iterate over all the numbers up to the limit
	for i := 2; i <= limit; i++ {
		// If the current number is not marked as composite
		if !composites[i] {
			// Mark all multiples of the current number as composite
			for j := i * 2; j <= limit; j += i {
				composites[j] = true
			}
			// Add the current number to the list of prime numbers
			primes = append(primes, i)
		}
	}
	return primes
}

func main() {
	limit := 2000000
	primes := sieveOfEratosthenes(limit)
	acc := 0
	for i := 0; i < len(primes); i++ {
		acc += primes[i]
	}
	fmt.Println(acc)
}
