package main

import (
	"fmt"
	"time"
)

type order struct {
	id        string
	amount    float32
	status    bool
	createdAt time.Time //nanosecond
}

// recieverType | added this func to order
func (o *order) changeStatus(status bool) {
	o.status = status
}

func (o *order) manipulateAmount() {
	o.amount *= 100
}

// constructor
func newOrder(id string, amount float32, status bool) *order {
	myOrder := order{
		id:        id,
		amount:    amount,
		status:    status,
		createdAt: time.Now(),
	}
	return &myOrder
}

func main() {
	// var o1 order = order{
	// 	id:     "001",
	// 	amount: 50.5,
	// 	status: true,
	// }

	o1 := order{
		id:     "001",
		amount: 50.5,
		status: true,
	}
	o1.createdAt = time.Now()
	fmt.Println(o1)

	o2 := order{
		id:        "002",
		amount:    150.5,
		status:    false,
		createdAt: time.Now(),
	}
	fmt.Println(o2)
	o2.changeStatus(true)
	o2.manipulateAmount()
	fmt.Println(o2)

	// what happens if we didnt give any particular val to var
	// it automatically get falsy value for it
	o3 := order{
		id:        "003",
		status:    false,
		createdAt: time.Now(),
	}
	fmt.Println(o3)

	o4 := newOrder("004", 10, false)
	fmt.Println(*o4)

	language := struct {
		name   string
		isGood bool
	}{"golang", true}

	fmt.Println(language)
}
