package main

import "fmt"

func handleTwoDigit(n int, lookup map[int]string) string {
	firstDigit := int(n/10) * 10
	secondDigit := n % 10
	if firstDigit == 0 {
		return lookup[secondDigit]
	}
	if firstDigit == 1 {
		return lookup[n]
	}
	if secondDigit == 0 {
		return lookup[firstDigit]
	}
	return fmt.Sprintf("%s-%s", lookup[firstDigit], lookup[secondDigit])
}

func main() {
	// two
	// forty-
	// one hundred and
	intToString := map[int]string{
		1:  "one",
		2:  "two",
		3:  "three",
		4:  "four",
		5:  "five",
		6:  "six",
		7:  "seven",
		8:  "eight",
		9:  "nine",
		10: "ten",
		11: "eleven",
		12: "twelve",
		13: "thirteen",
		14: "fourteen",
		15: "fifteen",
		16: "sixteen",
		17: "seventeen",
		18: "eighteen",
		19: "nineteen",
		20: "twenty",
		30: "thirty",
		40: "forty",
		50: "fifty",
		60: "sixty",
		70: "seventy",
		80: "eighty",
		90: "ninety",
	}

	for i := 1; i < 1000; i++ {
		if foo, ok := intToString[i]; ok {
			fmt.Println(foo)
		} else if i < 100 {
			fmt.Println(handleTwoDigit(i, intToString))
		} else {
			firstDigit := i / 100
			remainingDigits := i % 100
			if remainingDigits == 0 {
				fmt.Println(fmt.Sprintf("%s hundred", intToString[firstDigit]))
			} else {
				fmt.Println(fmt.Sprintf("%s hundred and %s", intToString[firstDigit], handleTwoDigit(remainingDigits, intToString)))
			}
		}
	}
	fmt.Println("one thousand")
}
