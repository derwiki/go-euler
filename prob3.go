package main
import (
    "fmt"
    "math"
)

// The prime factors of 13195 are 5, 7, 13 and 29.
//What is the largest prime factor of the number 600851475143 

func GeneratePrimes() chan int {
	out := make(chan int)
	go func() {
		primes := []int{2}
		for {
			for _, prime := range primes {
				out <- prime
			}
			var nextPrime int
			for {
				nextPrime += 2
				isPrime := true
				for _, prime := range primes {
					if nextPrime%prime == 0 {
						isPrime = false
						break
					}
				}
				if isPrime {
					break
				}
			}
			primes = append(primes, nextPrime)
		}
	}()
	return out
}

func SieveOfEratosthenes(limit int) []int {
	var primes []int
	sieve := make([]bool, limit+1)

	for i := 2; i <= limit; i++ {
		if !sieve[i] {
			primes = append(primes, i)
			for j := 2 * i; j <= limit; j += i {
				sieve[j] = true
			}
		}
	}

	return primes
}

func main() {
    target := 600851475143
    fmt.Println(target)
    sqrt := math.Sqrt(float64(target))
    fmt.Println(sqrt)
		primes := SieveOfEratosthenes(int(sqrt))
    for i := len(primes) - 1; i >= 0; i-- {
        if target % primes[i] == 0 {
            fmt.Println(primes[i])
            break
        }
		}
}
