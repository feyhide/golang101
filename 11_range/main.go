package main

import "fmt"

// iterating over data structures

func main() {
	nums := []int{1, 2, 3}

	// for i := 0; i < len(nums); i++ {
	// 	fmt.Println(nums[i])
	// }

	sum := 0
	for i, num := range nums {
		sum += num
		fmt.Println(i, num)
	}

	fmt.Println(sum)

	m := map[string]string{"name": "fahad", "uni": "fast"}

	for key, val := range m {
		fmt.Println(key, val)
	}

	for i, c := range "feyhide" {
		// i = index, c = uni code (rune)
		// i = starting byte of rune
		fmt.Println(i, c)

		fmt.Println(i, string(c))

	}
}
