package main

import "fmt"

// only for-loop for looping in go nothing else

func main() {
	i := 1

	for i <= 3 {
		fmt.Println(i)
		i++
	}

	// for {
	// 	fmt.Println("infinite")
	// }

	//classic forloop
	for i := 4; i < 8; i++ {
		if i == 6 {
			continue
		}

		if i == 7 {
			break
		}

		fmt.Println(i)
	}

	//range
	fmt.Println("Range ")
	for i := range 20 {
		fmt.Println(i)
	}
}
