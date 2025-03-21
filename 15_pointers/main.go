package main

import "fmt"

// pass by value
func passByVal(num int) {
	num = 5
	fmt.Println("In Func: ", num)
}

// pass by ref
func passByRef(num *int) {
	*num = 5
	fmt.Println("In Func: ", *num)
}

func main() {
	num := 1

	passByVal(num)

	fmt.Println("after passByVal in main ", num)

	passByRef(&num)

	fmt.Println("after passByRef in main ", num)
}
