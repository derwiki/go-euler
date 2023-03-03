package main

import "fmt"

func main() {
	/*
			43 44 45 46 47 48 49
			42 21 22 23 24 25 26
			41 20  7  8  9 10 27
			40 19  6  1  2 11 28
			39 18  5  4  3 12 29
			38 17 16 15 14 13 30
			37 36 35 34 33 32 31

			1x1 - center
			3x3 - TR 9, subtract 2: 9, 7, 5, 3
			5x5 - TR 25, subtract 4: 25, 21, 17, 13
			7x7 - TR 49, subtract 6: 49, 43, 37, 31

		    i := 1...1001 by twos {}
	*/

	acc := 1 // center square
	for i := 3; i <= 1001; i += 2 {
		acc += 4*i*i - (i-1)*6
	}
	fmt.Println("acc", acc)
}
