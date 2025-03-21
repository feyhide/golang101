package main

import (
	"fmt"
	"maps"
)

func main() {
	m := make(map[string]string)

	m["name"] = "feyhide"
	m["lang"] = "go"

	fmt.Println(m["name"], m["lang"])

	//if key doesnt exist it returns zero value
	fmt.Println(m["empty"])

	fmt.Println(len(m))

	delete(m, "lang")
	fmt.Println(m)

	clear(m)
	fmt.Println(m)

	m1 := map[string]int{"price": 40, "capacity": 10}
	m2 := map[string]int{"price": 40, "capacity": 10}
	fmt.Println(m2)

	k, ok := m2["price"]

	// k = value , ok is bool for found or not found
	fmt.Println(k)
	if ok {
		fmt.Println("all ok")
	} else {
		fmt.Println("not ok")
	}

	fmt.Println(maps.Equal(m1, m2))
}
