package helloworld

import (
  "fmt"
)

fun Ifthenelse() {
  var val int
  val=7
  if val%2 == 0 {
    fmt.Printf("The number %d is even\n", val)
  }else{
      fmt.Printf("The number %d is odd because remainder is %d\n",val,(val%2))
  }


}
