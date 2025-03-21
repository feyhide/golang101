package main

import "fmt"

const age = 40

//shorthand cannot be used outside

func main() {
	const name = "fahad"

	const (
		port = 3000
		host = "localhost"
	)

	fmt.Println(name)
	fmt.Println(age)
	fmt.Println(port, host)
}
