package main

import "fmt"

var y = 50 //variable declaration

func main() {
	x := 40

	fmt.Println(x)
	fmt.Println(y)

	var z int     // initialized to 0
	var f bool    // initialized to false
	var s string  // initialized to ""
	var d float64 // initialized to 0.0

	fmt.Println(z)
	fmt.Println(f)
	fmt.Println(s)
	fmt.Println(d)

	foo()
}

func foo() {
	//fmt.Print(x) // this doesnt work
	fmt.Print(y) // this does

}

// 40
// 50
// 0
// false

// 0
// 50
