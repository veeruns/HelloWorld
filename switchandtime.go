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
//Whatami returns type of the
func Whatami(i interface{}) string{
  switch i.(type) {
  case string:
      return "string"
  case int32,int64,int:
      return "integer"
  case bool:
      return "Boolean"
  default:
      return "UnknownType"
  }

}
