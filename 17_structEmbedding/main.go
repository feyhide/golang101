package main

import (
	"fmt"
	"time"
)

type customer struct {
	name  string
	phone string
}

type order struct {
	id        string
	amount    float32
	status    bool
	createdAt time.Time //nanosecond
	customer
}

func main() {
	// c1 := customer{
	// 	name:  "feyhide",
	// 	phone: "0900",
	// }

	// o1 := order{
	// 	id:       "001",
	// 	amount:   50.5,
	// 	status:   true,
	// 	customer: c1,
	// }

	o1 := order{
		id:     "001",
		amount: 50.5,
		status: true,
		customer: customer{
			name:  "feyhide",
			phone: "0900",
		},
	}

	o1.createdAt = time.Now()
	o1.customer.name = "notfeyhide"
	fmt.Println(o1)

}
