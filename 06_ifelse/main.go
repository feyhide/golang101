package main

import "fmt"

func main() {
	age := 50

	if age > 40 {
		fmt.Println("old as funk")
	} else if age >= 18 {
		fmt.Println("Adult")
	} else {
		fmt.Println("no Adult")
	}

	role := "admin"
	hasPermission := false

	if role == "admin" || hasPermission {
		fmt.Println("all in")
	}

	//declaring var in ifelse

	if age := 15; age > 18 {
		fmt.Println("again adult")
	} else if age <= 18 {
		fmt.Println("again not adult")
	}

	// go doesnot have ternary
}
