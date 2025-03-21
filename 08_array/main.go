package main

import (
	"fmt"
)

// we use array in go if we already know size of array else we use slice

func main() {
	// auto add 0th/falsy values
	var arr [5]int
	fmt.Println(arr)

	var vals [4]bool
	fmt.Println(vals)

	var str [4]string
	fmt.Println(str)

	fmt.Println(len(arr))
	arr[0] = 1
	fmt.Println(arr[0])

	//declaring in single line
	arr1 := [3]int{1, 2, 4}
	fmt.Println(arr1)

	arr2d := [2][2]int{{1, 3}, {2, 4}}
	fmt.Println(arr2d)
}
