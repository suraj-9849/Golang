package main

import "fmt"

// THis is not the correct way of doing because we are not following the Open-Close Principle we are acutally voilating it so lets see the better approach

// type payment struct{}
// func (p payment) makePayment(amount float32) {
// 	// razorpayPW:=razorpay{}
// 	// razorpayPW.pay(amount)
// 	stripePW:=stripe{}
// 	stripePW.pay(amount)
// }

// type razorpay struct {}

// func (r razorpay) pay(amount float32){
// 	fmt.Println("Making Payment using RazorPay:", amount)
// }

// type stripe struct {}

// func (s stripe) pay(amount float32){
// 	fmt.Println("Making Payment using Stripe:", amount)
// }

type paymenter interface{
	pay(amount float32)
}

type payment struct{
	gateWay paymenter  
}
func (p payment) makePayment(amount float32) {
	if p.gateWay == nil {
		fmt.Println("No payment gateway provided")
		return
	}
	p.gateWay.pay(amount)
}

type razorpay struct {}

func (r razorpay) pay(amount float32){
	fmt.Println("Making Payment using RazorPay:", amount)
}

type stripe struct {}

func (s stripe) pay(amount float32){
	fmt.Println("Making Payment using Stripe:", amount)
}
func main() {
	newPayment := payment{gateWay: stripe{}}
	newPayment.makePayment(56.85)
}