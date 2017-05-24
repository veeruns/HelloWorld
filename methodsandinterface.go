package helloworld

import (
"fmt"
)

type Circle struct {
  var x float64
  var y float64
  var r float64
}

func Circlestuff(){
  var c Circle
  x := new(Circle)
  c.x=10
  c.y=20
  c.r=30
  fmt.Printf("%f %f %f\n",x.x,x.y,x.r)
  fmt.Printf("%f %f %f\n",c.x,c.y,c.r)
}
