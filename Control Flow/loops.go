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
