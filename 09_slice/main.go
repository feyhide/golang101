package main

import (
	"fmt"
	"slices"
)

//slice -> dynamic

func main() {
	//uninitialized slice is nil
	var nums []int
	fmt.Println(nums == nil)

	//(type,initial size, capacity)
	var num2 = make([]int, 1, 3)
	// initial size is 2

	fmt.Println(num2)
	fmt.Println(num2 == nil)
	//capacity auto increases
	fmt.Println(cap(num2))

	num2[0] = 5
	num2 = append(num2, 2)
	num2 = append(num2, 3)
	num2 = append(num2, 4)
	fmt.Println(num2)
	fmt.Println(cap(num2))

	num3 := []int{}
	fmt.Println(num3)
	fmt.Println(cap(num3))
	num3 = append(num3, 2)
	fmt.Println(num3)
	fmt.Println(cap(num3))

	//copy func
	var num4 = make([]int, len(num2))
	fmt.Println(num4)
	copy(num4, num2)
	fmt.Println(num4)

	//slice operator
	var num5 = []int{1, 2, 3}

	fmt.Println(num5[0:2])

	//slice package
	var arw1 = []int{1, 2}
	var arw2 = []int{1, 2}

	fmt.Println(slices.Equal(arw1, arw2))

	//2d slices
	var _2dslice = [][]int{{1, 2, 3}, {4, 5, 6}}
	fmt.Println(_2dslice)
}
