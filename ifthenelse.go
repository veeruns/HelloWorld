package helloworld

import (
	"fmt"
)

// Ifthenelse function prints a value is even or odd
func Ifthenelse() {
	var val int
	val = 7
	if val%2 == 0 {
		fmt.Printf("The number %d is even\n", val)
	} else {
		fmt.Printf("The number %d is odd because remainder is %d\n", val, (val % 2))
	}

}

//Whatnumber function accepts a number as input and returns in string if number is negative, positive or zero
func Whatnumber(input int) string {
	if input < 0 {
		return "negative"
	} else if input > 0 {
		return "positive"
	} else {
		return "Zero"
	}
}
