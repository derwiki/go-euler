package main

import "fmt"

func findDivisors(n int) []int {
	divisors := []int{}
	for i := 1; i <= (n/2)+1; i++ {
		if n%i == 0 {
			divisors = append(divisors, i)
		}
	}
	return divisors
}
func sumList(xs []int) int {
	acc := 0
	for i := 0; i < len(xs); i++ {
		acc += xs[i]
	}
	return acc
}
func main() {
	acc := 0
	for i := 0; i <= 10_000; i++ {
		a := sumList(findDivisors(i))
		if sumList(findDivisors(a)) == i && i != a {
			fmt.Println(fmt.Sprintf("Amicable: %d %d", i, a))
			acc += i
		}
	}
	fmt.Println(acc)
}
