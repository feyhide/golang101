package main

import "fmt"

func sum(nums ...int) int {
	//nums gets us a slice
	total := 0
	for _, nums := range nums {
		total += nums
	}

	return total
}

func main() {
	// n parameters can be passed
	//fmt.Println(1, 2, 3, 4)

	var result int = sum(4, 5, 6, 7)
	fmt.Println(result)

	nums := []int{1, 2, 3}
	result = sum(nums...)
	fmt.Println(result)
}
