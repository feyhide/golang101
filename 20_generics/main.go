package main

import "fmt"

// allow all [T any]
// bounded as below
// func printSlice[T int | string](items []T) {
// 	for _, item := range items {
// 		fmt.Println(item)
// 	}
// }

// allow all comparable types
// func printSlice[T comparable](items []T) {
// 	for _, item := range items {
// 		fmt.Println(item)
// 	}
// }

// multiple
func printSlice[T comparable, V string](items []T, uni V) {
	for _, item := range items {
		fmt.Println(item, uni)
	}
}

type stack[T any] struct {
	elements []T
}

func main() {
	nums := []int{1, 2, 3}
	str := []string{"fahad", "kazim"}
	printSlice(str, "fast")
	printSlice(nums, "nuces")

	s1 := stack[int]{
		elements: []int{1, 2, 3},
	}

	s2 := stack[string]{
		elements: []string{"fahad", "kazim"},
	}

	fmt.Println(s1)
	fmt.Println(s2)
}
