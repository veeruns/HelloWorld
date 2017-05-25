package main

import (
	"fmt"

	"github.com/veeruns/codekata"
	"github.com/veeruns/helloworld"
)

func main() {
	helloworld.Arraystuff()
	helloworld.Hashprint()
	helloworld.Ifthenelse()
	s := helloworld.Whatnumber(0)
	fmt.Printf("Returned %s\n", s)
	x := []int{1, 2, 3, 4, 5}
	i := helloworld.Addarray(x)
	fmt.Printf("Addition of array %d\n", i)
	w := helloworld.WhatDay()
	fmt.Printf("Its a %s\n", w)
	q := helloworld.Whatami(true)
	fmt.Printf("Whatami returns %s\n", q)

	fmt.Printf("Whatami returns %s\n", helloworld.Whatami(5))
	fmt.Printf("Whatami returns %s\n", helloworld.Whatami("test"))
	num := helloworld.Addnumbers(1, 2, 3)
	fmt.Printf("Addition is %d\n", num)
	helloworld.Circlestuff()
	helloworld.PrintRandom()
	codekata.Flagoperation()
	var stuff string
	stuff = codekata.Logfile
	timeout := codekata.Timeout
	fmt.Printf("Log file name is %s\n", stuff)
	fmt.Printf("Timeout value is %d\n", timeout)
}
