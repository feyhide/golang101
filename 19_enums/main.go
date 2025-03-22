package main

import "fmt"

type OrderStatus int

const (
	Recieved OrderStatus = iota
	Confirmed
	Prepared
	Delivered
)

func changeOrderStatus(status OrderStatus) {
	fmt.Println("Changing order status to", status)
}

type OrderStatus2 string

const (
	Recieved2  OrderStatus2 = "recieved"
	Confirmed2 OrderStatus2 = "confirmed"
	Prepared2  OrderStatus2 = "prepared"
	Delivered2 OrderStatus2 = "delivered"
)

func changeOrderStatus2(status OrderStatus2) {
	fmt.Println("Changing order status to", status)
}

func main() {
	changeOrderStatus(Confirmed)
	changeOrderStatus2(Confirmed2)
}
