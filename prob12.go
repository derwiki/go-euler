package main

import "fmt"

func divisorCount(a int) int {
	acc := 1
	for i := 1; i <= a/2; i++ {
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
		if divisors > 500 {
			fmt.Println(fmt.Sprintf("Final: %d: %d - %d", i, triangle, divisors))
			return
		}
	}
}
