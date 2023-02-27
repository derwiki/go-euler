package main

import "fmt"

func findDivisors2(n int) []int {
	divisors := []int{}
	for i := 1; i <= (n/2)+1; i++ {
		if n%i == 0 {
			divisors = append(divisors, i)
		}
	}
	return divisors
}
func sumList2(xs []int) int {
	acc := 0
	for i := 0; i < len(xs); i++ {
		acc += xs[i]
	}
	return acc
}

func main() {
	abunduntNumbersMap := make(map[int]bool)
	var abunduntNumbers []int
	for i := 12; i <= 28123; i++ {
		if sumList2(findDivisors2(i)) > i {
			abunduntNumbersMap[i] = true
			abunduntNumbers = append(abunduntNumbers, i)
		}
	}
	acc := 0
	for i := 0; i <= 28123; i++ {
		canBeWrittenAsSum := false
		for _, abunduntNumber := range abunduntNumbers {
			if abunduntNumbersMap[i-abunduntNumber] {
				canBeWrittenAsSum = true
				break
			}
		}

		if !canBeWrittenAsSum {
			acc += i
		}
	}
	fmt.Println("acc", acc)
}
