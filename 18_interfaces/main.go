package main

import "fmt"

// type payment struct{}

//OPEN CLOSE PRINCIPLE VIOLATED
// func (p payment) makePayment(amount float32) {
// 	//stripePaymentGateway := stripe{}
// 	razorpayPaymentGateway := razorpay{}
// 	//stripePaymentGateway.pay(amount)
// 	razorpayPaymentGateway.pay(amount)
// }

type paymenter interface {
	pay(amount float32)
}

type payment struct {
	gateway paymenter
}

func (p payment) makePayment(amount float32) {
	p.gateway.pay(amount)
}

type stripe struct{}

func (s stripe) pay(amount float32) {
	fmt.Println("making payment using stripe", amount)
}

type razorpay struct{}

func (s razorpay) pay(amount float32) {
	fmt.Println("making payment using razorpay", amount)
}

func main() {
	p1 := payment{gateway: stripe{}}
	p1.makePayment(100)

	p2 := payment{gateway: razorpay{}}
	p2.makePayment(200)
}
