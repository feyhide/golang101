package main

import "fmt"

func main() {
	var choice int

	fmt.Println("Choose an example to run:")
	fmt.Println("1 - Example 1")
	fmt.Println("2 - Example 2")
	fmt.Println("3 - ex3: web api")
	fmt.Print("Enter your choice (1/2/3): ")

	fmt.Scanln(&choice)

	switch choice {
	case 1:
		runEx1()
	case 2:
		runEx2()
	case 3:
		runEx3()
	default:
		fmt.Println("Invalid choice! Please enter 1, 2, or 3.")
	}
}
