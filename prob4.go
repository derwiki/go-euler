package main
import (
  "fmt"
  "strconv"
)

func isFib(x int, y int) bool {
  buffer := strconv.Itoa(x * y)
  midpoint := len(buffer) / 2
  //fmt.Println(midpoint)
  for i := 0; i < midpoint; i++ {
      //fmt.Println(i)
      //fmt.Println(len(buffer) - 1 - i)
      //fmt.Println(buffer[len(buffer) - 1 - i])
      //fmt.Println("*****")
      if buffer[len(buffer) - 1 - i] != buffer[i] {
          return false
      }
      //fmt.Sprintf("%d\t%d", buffer[i], buffer[len(buffer) - i])
  }
  return true
}

func main() {
    x := 999
    y := 999

    for ;x>0; x--{
       for ;y>0; y--{
           if isFib(x, y) {
               fmt.Println(fmt.Sprintf("%d, %d, %t", x, y, isFib(x, y)))
           }
       }
    }
    //fmt.Println(isFib(x, y))
    //fmt.Println(isFib(91, 99))

}
