package main

import "fmt"

func main() {
	var name string = "feyhide"
	var isLive bool = true
	var age int = 21
	var price float32 = 30.5

	//infer
	var name1 = "feyhide"
	var isLive1 = false
	var age1 = 21
	var price1 = 21.5

	//shorthand syntax
	uni := "fast nuces"

	// we cant really use shorthand if were going to assign var later
	// then we have to use type too

	var later string

	fmt.Println(name)
	fmt.Println(name1)
	fmt.Println(isLive)
	fmt.Println(isLive1)
	fmt.Println(age)
	fmt.Println(age1)
	fmt.Println(price)
	fmt.Println(price1)
	fmt.Println(uni)
	later = "later gator"
	fmt.Println(later)
}
