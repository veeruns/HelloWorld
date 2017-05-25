package helloworld

import (
"fmt"
)

type Circle struct {
  x float64
  y float64
  r float64
}

func Circlestuff(){
  var c Circle
  x := new(Circle)
  c.x=10
  c.y=20
  c.r=30
  // new initializes the data to zero
  fmt.Printf("%f %f %f\n",x.x,x.y,x.r)
  fmt.Printf("%f %f %f\n",c.x,c.y,c.r)
}
