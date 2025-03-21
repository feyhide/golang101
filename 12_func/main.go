package main

import "fmt"

func add(a, b int) int {
	return a + b
}

func getlang() (string, string, bool) {
	return "golang", "js", true
}

func processIt(fn func(a int) int) int {
	return fn(1)
}

func main() {
	result := add(6, 9)
	fmt.Println(result)

	fmt.Println(getlang())
	lang1, lang2, status := getlang()
	fmt.Println(lang1, lang2, status)

	fn := func(a int) int {
		return a + 2
	}

	idk := processIt(fn)
	fmt.Println(idk)
}
