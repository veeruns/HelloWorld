package helloworld
import (
"fmt"
)

func Addnumbers(a ...int) int {
  sum := 0
  fmt.Printf("Addnumber function\n")
  for _,val := range a {
    sum += val
  }
  return sum
}
