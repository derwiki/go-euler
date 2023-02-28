package main

import (
	"fmt"
	"math/big"
)

func FindLongestRecurringCycle(s string) string {
	n := len(s)
	maxCycle := 0
	maxCycleStr := ""

	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			k := i
			for l := j; l < n && s[k] == s[l]; l, k = l+1, k+1 {
				if l-j+1 > maxCycle {
					maxCycle = l - j + 1
					maxCycleStr = s[j : j+maxCycle]
				}
			}
		}
	}

	return maxCycleStr
}

func findRepeatedCycles(s string) string {
	longest := ""
	// skip the first j characters
	for j := 0; j < len(s); j++ {
		for length := 1; j+length < len(s); length++ {
			searchTerm := s[j : j+length]
			fmt.Println(j, length, searchTerm, len(s)-j-length)
			// TODO(derwiki) need to tighten up j < length
			for k := 1; j < length; k++ {
				offset := len(searchTerm) * k
				nextPossibleMatch := s[j+offset : j+offset+len(searchTerm)]
				fmt.Println(nextPossibleMatch)
				if nextPossibleMatch == searchTerm {
					fmt.Println("found cycle")
				}
			}
		}
	}
	return longest
}

func findRepeatedCycles2(s string) string {
	longest := ""
	for i := 0; i < len(s); i++ {
		for length := 1; length < len(s)/2+1; length++ {
			chunks := splitString(s, length)
			searchChunk := chunks[0]
			foundCycles := 0
			for _, chunk := range chunks[1:] {
				fmt.Println(i, length, searchChunk, chunk)
				if searchChunk == chunk {
					foundCycles++
				}
			}
			fmt.Println("found cycles", foundCycles)
			if foundCycles == len(chunks)-1 {
				fmt.Println("found repeating cycle")
			}
		}
	}
	return longest
}

func splitString(s string, chunkSize int) []string {
	var chunks []string

	for i := 0; i < len(s); i += chunkSize {
		end := i + chunkSize
		if end > len(s) {
			end = len(s)
		}
		chunks = append(chunks, s[i:end])
	}

	return chunks
}

func main() {
	one := new(big.Float).SetPrec(1000).SetFloat64(1.0)
	for i := 990.0; i < 991.0; i++ {
		inverse := new(big.Float).SetPrec(1000).Quo(one, new(big.Float).SetPrec(1000).SetFloat64(float64((i))))
		fmt.Println(i, inverse)
		inverseString := inverse.Text('f', -1)
		fmt.Println(i, inverseString)
		choppedString := inverseString[2 : len(inverseString)-1]
		fmt.Println(i, choppedString)
		fmt.Println()
		findRepeatedCycles2("appleappleappleapple")
		// findRepeatedCycles("55001001001001001001001")
	}
}
