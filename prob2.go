package main
import "fmt"

func main() {
    fib0 := 0
    fib1 := 1
    evens := 0

    for ; fib1 < 4000000; {
        newfib := fib0 + fib1
        fib0 = fib1
        fib1 = newfib
        if newfib % 2 == 0 {
            evens += newfib
        }
    }
    fmt.Println(evens)
}
