package main

import "fmt"

func main() {

	var x [5]int //array
	for i := 0; i < 5; i++ {
		x[i] = i
	}
	fmt.Println(x)

	var y = []int{1, 2, 3, 4, 5} //slice
	fmt.Println(y)
	fmt.Println(len(y))
	// for i, v := range y {
	// 	fmt.Println(i, " --> ", v)
	// }
	var z []int
	z = append(y, y...)
	fmt.Println(z) // z becomes [1 2 3 4 5 1 2 3 4 5]

}

// [0 1 2 3 4]
// [1 2 3 4 5]
// 5
// [1 2 3 4 5 1 2 3 4 5]
