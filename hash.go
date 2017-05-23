package helloworld

import (
	"fmt"
)

// Hashprint prints a hash
func Hashprint() {

	hash := make(map[string]int)
	hash["t1"] = 5
	hash["t2"] = 6
	hash["t3"] = 7
	for k, v := range hash {
		fmt.Printf("%s,%d\n", k, v)
	}
}
