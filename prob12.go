package main

import "fmt"

func divisorCount(a int) int {
	acc := 0
	for i := 1; i <= a; i++ {
		if a%i == 0 {
			acc += 1
		}
	}
	return acc
}

func main() {
	triangle := 0
	for i := 1; ; i++ {
		triangle += i
		divisors := divisorCount(triangle)
		fmt.Println(fmt.Sprintf("%d: %d - %d", i, triangle, divisors))
		if divisors > 500 {
			return
		}
	}
}
