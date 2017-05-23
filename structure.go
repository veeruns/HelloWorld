package helloworld

import (
 "fmt"
)
type Employee struct {
  firstname string
  lastname string
  address Address
}

type Address struct {
  Street string
  City string
  State string
  zipcode string
  func PrintAddress () {
    fmt.Printf("%s,%s\n%s\n",Street,City,zipcode)
  }
}
