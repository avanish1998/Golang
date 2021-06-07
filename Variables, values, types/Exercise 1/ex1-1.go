/*Assign values to variables 'x' , 'y' , 'z'
42,
James Bond,
true

Print values using
single print
multiple print

Use sprintf to assign all these values to a variable s
*/

package main

import (
	"fmt"
)

var x = 42
var y = "James Bond"
var z = true

func main() {
	fmt.Println(x, y, z)
	fmt.Println(x)
	fmt.Println(y)
	fmt.Println(z)
	s := fmt.Sprintf("%v\t%v\t%v", x, y, z)
	fmt.Println(s)

}
