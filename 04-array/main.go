// Arrays is a group of data/values with same type
// Arrays in go have preset size and can not expand

package main

import (
	"fmt"
)

func main() {
	names := [3]string{"Ahmed", "Ali", "Khaled"}
	fmt.Println(len(names))
	fmt.Println(names[len(names)-1])

	xoBoard := [3][3]string{
		[3]string{"-", "-", "x"},
		[3]string{"-", "o", "-"},
		[3]string{"x", "o", "-"}, // note we need to add that , to enable new line
	}

	fmt.Println(xoBoard)

	xoBoard[0][1] = "x"

	fmt.Println(xoBoard)

}
