package main
import "fmt"

func main() {
  for i:= 20;; i++ {
	  foo := true
		for j := 1; j <= 20; j++ {
		  if i % j != 0 {
			  foo = false
		  }
		}
		if foo == true {
		  fmt.Println(i)
			return
		}
	}
}
