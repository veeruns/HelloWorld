package helloworld

import (
	"fmt"
	"strings"
)

func foo(hello, world string) (v, w string) {
	v = strings.Join([]string{hello, world}, " ")
	w = strings.Join([]string{hello, world}, "+")
	return v, w
}

// Freturn experiments with function returning two values
func Freturn() {
	var r1, r2 string
	r1, r2 = foo("Hello", "World")
	fmt.Printf("%s\n", r1)
	fmt.Printf("%s\n", r2)
}
