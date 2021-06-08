package main

import "fmt"

func main() {
	for i := 0; ; i++ {
		if i < 5 {
			continue
		} else if i > 10 {
			break
		} else {
			fmt.Println(i)
		}
	}

}

// 5
// 6
// 7
// 8
// 9
// 10
