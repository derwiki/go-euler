package main

import "fmt"

func collatzSequenceLength(n int) int {
	acc := 1
	for n > 1 {
		if n%2 == 0 {
			n /= 2
		} else {
			n *= 3
			n++
		}
		acc++
	}
	return acc
}

func main() {
	longestCount := 0
	longestCountStart := 0
	for i := 2; i < 1000000; i++ {
		count := collatzSequenceLength(i)
		if count > longestCount {
			longestCount = count
			longestCountStart = i
		}
	}
	fmt.Println(fmt.Sprintf("longestCount: %d \n longestCountStart: %d", longestCount, longestCountStart))
}
