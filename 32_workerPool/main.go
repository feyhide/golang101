package main

import "fmt"

func main() {
	var choice int

	fmt.Println("Choose an example to run:")
	fmt.Println("1 - basic worker pool with 5 parallel workers")
	fmt.Println("2 - worker pool with diff type of tasks with 5 parallel workers")
	fmt.Print("Enter your choice (1/2/3): ")

	fmt.Scanln(&choice)

	switch choice {
	case 1:
		singleType()
	case 2:
		diffTypes()
	default:
		fmt.Println("Invalid choice! Please enter 1, 2, or 3.")
	}
}
