package main

import "fmt"

//enum-> enumerated Types:

type OrderStatus int

const (
	Recieved OrderStatus = iota
	Delivered
	Failed
	Confirmed
)

func changeOrderStatus(status OrderStatus){
	fmt.Println("Changing the order status to: ",status)
}

func main(){
	changeOrderStatus(Delivered)
}