package main

import "fmt"

var x int

type avanish int // creating type avanish
var y avanish

func main() {
	x = 40
	y = 60
	fmt.Println(x)
	fmt.Printf("%T\n", x)
	fmt.Println(y)
	fmt.Printf("%T\n", y)
	// b = a  -->  doesnt work
}

// 40
// int
// 60
// main.avanish
