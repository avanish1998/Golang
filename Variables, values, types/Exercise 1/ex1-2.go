package main

import "fmt"

type ag int

var x ag

func main() {
	fmt.Println(x)
	fmt.Printf("%T\n", x)
	x = 40
	fmt.Println(x)
}
