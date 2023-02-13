package main
import "fmt"

func main() {
  sum_of_squares := 0
	sums := 0
  for i := 1; i <= 100; i++ {
	  sum_of_squares += i * i
		sums += i
	}
	fmt.Println(sums * sums - sum_of_squares)
}
