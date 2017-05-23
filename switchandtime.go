package helloworld
import (
 "fmt"
 "time"
)

func WhatDay()string {
  switch time.Now().Weekday() {
  case time.Saturday,time.Sunday :
      return "Weekend"
  default:
      return "Weekday"
  }
}
