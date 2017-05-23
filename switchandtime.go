package helloworld
import (
 "fmt"
 "time"
)
//WhatDay returns Whatday it is weekday or weekend
func WhatDay()string {
  fmt.Printf("Testing Testing Testing")
  switch time.Now().Weekday() {
  case time.Saturday,time.Sunday :
      return "Weekend"
  default:
      return "Weekday"
  }
}
