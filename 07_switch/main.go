package main

import (
	"fmt"
	"time"
)

func main() {
	i := 2

	// no need to write break
	switch i {
	case 1:
		fmt.Println("one")
	case 2:
		fmt.Println("two")
	case 3:
		fmt.Println("three")
	case 4:
		fmt.Println("four")
	case 5:
		fmt.Println("five")
	default:
		fmt.Println("Not in range 1-5")
	}

	//multiple cond switch
	switch time.Now().Weekday() {
	case time.Saturday, time.Sunday:
		//working as or
		fmt.Println("its weekend")
	default:
		fmt.Println("Work day")
	}

	hour := time.Now().Hour()

	switch {
	case hour >= 9 && hour <= 17:
		fmt.Println("Work hours")
	default:
		fmt.Println("Off hours")
	}

	whoAmI := func(i any) {
		// i.(type) returns type of i
		switch t := i.(type) {
		case int:
			fmt.Println("its an integer")
		case string:
			fmt.Println("its a string")
		default:
			fmt.Println("other :", t)
		}
	}

	whoAmI("fahad")
}
