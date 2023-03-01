package main

import (
	"fmt"
	"math/big"
)

func findRepeatedCycles(s string) string {
	// TODO cache of longest prefixes to bail out earlier
	longest := ""
	for length := 1; length < len(s); length++ {
		chunks := splitString(s, length)
		searchChunk := chunks[0]
		foundCycles := 0
		trimmedChunks := chunks[1 : len(chunks)-1]
		for _, chunk := range trimmedChunks {
			if searchChunk == chunk {
				foundCycles++
				// fmt.Println("found cycle", searchChunk, chunk)
				// TODO(derwiki) again does this need to be minus _two_
				if foundCycles == len(chunks)-1 || foundCycles == len(chunks)-2 {
					//fmt.Println("found repeating cycle", searchChunk, "**", chunk, "**", trimmedChunks)
					//fmt.Println("return with best cycle, match count:", foundCycles)
					return searchChunk
				}
			} else {
				//fmt.Println("found break in cycle", searchChunk, chunk)
				break
				// if the next chunk doesn't support the cycle, we can bail early
			}
		}
	}
	return longest
}
func findRepeatedCyclesWithSkips(s string) string {
	longest := ""
	for i := 0; i < len(s); i++ {
		cycle := findRepeatedCycles(s[i:])
		if len(cycle) > len(longest) {
			longest = cycle
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
	var precision uint = 10000
	one := new(big.Float).SetPrec(precision).SetFloat64(1.0)
	longestD := 0.0
	longestDLength := 0
	longestDString := ""
	// TODO(derwiki) why does this panic if we start at i=1?
	for i := 2.0; i <= 999.0; i++ {
		inverse := new(big.Float).SetPrec(precision).Quo(one, new(big.Float).SetPrec(precision).SetFloat64(float64((i))))
		//fmt.Println(i, inverse)
		inverseString := inverse.Text('f', -1)
		//fmt.Println(i, inverseString)
		choppedString := inverseString[2 : len(inverseString)-1]
		//fmt.Println(i, choppedString)
		//fmt.Println()
		result := findRepeatedCyclesWithSkips(choppedString)
		if len(result) > longestDLength {
			longestDLength = len(result)
			longestD = i
			longestDString = result
			fmt.Println("new longestD", longestD, "longestDLength", longestDLength, "cycle", longestDString)
		}
	}
	fmt.Println("overall longestD", longestD, "longestDLength", longestDLength, "cycle", longestDString)

	for _, v := range [][2]string{
		{"appleappleappleapple", "apple"},
		{"aappleappleappleapple", "apple"},
		{"55001001001001001001001", "001"},
		{"00100502512562814070351758793969849246231155778894472361809045226130653266331658291457286432160804020100502512562814070351758793969849246231155778894472361809045226130653266331658291457286432160804020100502512562814070351758793969849246231155778894472361809045226130653266331658291457286432160804020100", "010050251256281407035175879396984924623115577889447236180904522613065326633165829145728643216080402"},
	} {
		input, cycle := v[0], v[1]
		result := findRepeatedCyclesWithSkips(input)
		fmt.Println("input", input, "expected", cycle, "result:", result)
		if result == cycle {
			fmt.Println("PASSED")
		} else {
			fmt.Println("! FAILED")
		}
	}
}
