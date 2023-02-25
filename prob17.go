package main

import "fmt"

func handleTwoDigit(n int, lookup map[int]string) string {
	if n > 10 && n < 20 {
		return lookup[n]
	}
	firstDigit := int(n/10) * 10
	secondDigit := n % 10
	if firstDigit == 0 {
		return lookup[secondDigit]
	}
	if secondDigit == 0 {
		return lookup[firstDigit]
	}
	return fmt.Sprintf("%s%s", lookup[firstDigit], lookup[secondDigit])
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

	acc := 0
	var numString string
	for i := 1; i < 1000; i++ {
		if foo, ok := intToString[i]; ok {
			numString = foo
		} else if i < 100 {
			numString = handleTwoDigit(i, intToString)
		} else {
			firstDigit := i / 100
			remainingDigits := i % 100
			if remainingDigits == 0 {
				numString = fmt.Sprintf("%shundred", intToString[firstDigit])
			} else {
				numString = fmt.Sprintf("%shundredand%s", intToString[firstDigit], handleTwoDigit(remainingDigits, intToString))
			}
		}
		fmt.Println(numString)
		acc += len(numString)
	}
	numString = "onethousand"
	fmt.Println(numString)
	acc += len(numString)
	fmt.Println(acc)

}
