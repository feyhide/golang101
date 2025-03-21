package main

import "fmt"

func counter() func() int {
	var count int = 0

	return func() int {
		count++
		return count
	}
}

// A closure is a function that remembers variables from its outer scope,
// even after the outer function has finished executing.

func main() {
	inc := counter()

	fmt.Println(inc())
	fmt.Println(inc())
	fmt.Println(inc())
}
