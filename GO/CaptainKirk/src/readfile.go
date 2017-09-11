package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {

	// open a file
	if file, err := os.Open("test.txt"); err == nil {

		// make sure it gets closed
		defer file.Close()

		// create a new scanner and read the file line by line
		scanner := bufio.NewScanner(file)
		var count int
		var line_to_array [][]byte
		for scanner.Scan() {

			if count == 0 {
				for i:=0; i < 4 ; i ++ {
            for j:=0;j < len(line_to_array); j++ {
              line_to_array[i][j]=0
            }
        }
			}

			arraybuffer := []byte(scanner.Text())
			copy(line_to_array[count], arraybuffer)
			for i := 0; i < 4; i++ {
				for j := 0; j < len(line_to_array[i]); j++ {
					fmt.Printf("%c", line_to_array[i][j])
				}
				fmt.Printf("\n")
			}

			count++
			if count == 3 {
				count = 0
			}

		}

		// check for errors
		if err = scanner.Err(); err != nil {
			log.Fatal(err)
		}

	} else {
		log.Fatal(err)
	}

}
