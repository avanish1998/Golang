package main

import "fmt"

func main() {
	var x = "Strings in Golang."
	fmt.Println(x)
	fmt.Printf("%T\n", x)

	for i := 0; i < len(x); i++ {
		fmt.Printf("%c", x[i])
	}
	fmt.Println()
	for i, v := range x {
		fmt.Printf("%d\t%c\n", i, v)
	}

}

// Strings in Golang.
// string
// Strings in Golang.
// 0       S
// 1       t
// 2       r
// 3       i
// 4       n
// 5       g
// 6       s
// 7
// 8       i
// 9       n
// 10
// 11      G
// 12      o
// 13      l
// 14      a
// 15      n
// 16      g
// 17      .
