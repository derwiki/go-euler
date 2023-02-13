package main

import (
	"fmt"
	"strconv"
)

func isFib(x int, y int) bool {
	buffer := strconv.Itoa(x * y)
	midpoint := len(buffer) / 2
	for i := 0; i < midpoint; i++ {
		if buffer[len(buffer)-1-i] != buffer[i] {
			return false
		}
	}
	return true
}

func main() {
	largestPalindrome := 0
	for x := 999; x > 99; x-- {
		for y := 999; y > 99; y-- {
			if isFib(x, y) {
				fmt.Println(fmt.Sprintf("%d, %d, %d, %t", x, y, x*y, isFib(x, y)))
				if x*y > largestPalindrome {
					largestPalindrome = x * y
				}
			}
		}
	}
	fmt.Println(fmt.Sprintf("Largest: %d", largestPalindrome))
}
