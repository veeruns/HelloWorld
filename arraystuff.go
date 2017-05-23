package helloworld

import (
	"fmt"
)

func arrayaddition([]int) int {
	return 0
}

// Arraystuff prints an array
func Arraystuff() {
	x := []int{1, 2, 3, 4}
	x = append(x, 5)
	for i := 0; i < len(x); i++ {
		fmt.Printf("%d\n", x[i])
	}
	for idx, y := range x {
		fmt.Printf("%d,%d\n", idx, y)
	}
}

// Addarray will add all the numbers in the array and returns the sum
func Addarray(input []int) int {
	output := 0
	for _, y := range input {
		output += y
	}
	return output
}
