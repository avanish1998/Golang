package main

import "fmt"

func main() {
	var x bool
	var y = 10
	var z = 20
	x = (y > z)
	fmt.Println(x)
	x = (y < z)
	fmt.Print(x)
}

// false
// true
