package main

import "fmt"

var x int8  // -128 to 127
var y int16 // -32768 to 32767
func main() {
	x = 5
	// y = x  --> cannot use x (type int8) as type int16 in assignment
	y = int16(x) // --> this works, conversion not casting
	fmt.Println(y)
	y += 5
	// x = 128 --> error, constant 128 overflows int8
	x = int8(y) // conversion not casting
	fmt.Println(y)
}
