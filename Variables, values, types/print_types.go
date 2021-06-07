package main

import (
	"fmt"
)

func main() {
	var x = 22
	fmt.Println(x)
	fmt.Printf("%T\n", x) // prints type of x

	var y = "hello"
	fmt.Println(y)
	fmt.Printf("%T\n", y) // prints type of y
	//y = 43   --->  cannot use 43 (type untyped int) as type string in assignment

}

// 22
// int
// hello
// string
