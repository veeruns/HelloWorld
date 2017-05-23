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

func Whatami(i interface{}) string{
  switch t:=i.(type) {
  case string:
      return "string"
  case int32,int64:
      return "Integer"
  default:
      return "UnknownType"
  }

}
