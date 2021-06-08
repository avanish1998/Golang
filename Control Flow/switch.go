package main

import "fmt"

func main() {
	var x int
	x = 2
	switch {
	case (x == 1):
		fmt.Println("x == 1")
	case (x == 2):
		fmt.Println("x == 2")
		fallthrough
	case (x == 3):
		fmt.Println("x != 3")

	}
}

// x == 2
// x != 3
