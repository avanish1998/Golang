package main

import "fmt"

func main() {
	for j := 0; j < 10; j++ {
		for i := 0; i < j; i++ {
			fmt.Print(i, " ")
		}
		fmt.Println()
	}
}

// 0
// 0 1
// 0 1 2
// 0 1 2 3
// 0 1 2 3 4
// 0 1 2 3 4 5
// 0 1 2 3 4 5 6
// 0 1 2 3 4 5 6 7
// 0 1 2 3 4 5 6 7 8
