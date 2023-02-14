package main

import "fmt"

func main() {
	for c := 1; c < 1000; c++ {
		for b := 1; b < c; b++ {
			for a := 1; a < b; a++ {
				if a*a+b*b == c*c {
					if a+b+c == 1000 {
						fmt.Println(fmt.Sprintf("%d %d %d: %d", a, b, c, a*b*c))
						return
					}
				}
			}
		}
	}
}
