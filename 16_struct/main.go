package main

import (
	"fmt"
	"time"
)

type customer struct{
	name string
	phoneNum int
}

type Order struct{
	id string
	amount float32
	status string
	createdAt time.Time // precession is nano seconds
	// struct embedding
	customer
}

// Reciever Type:
func (o *Order) changeStatus(status string){
	o.status = status
}

func newOrder(id string, amount float32, status string, createdAt time.Time) *Order{
	myOrder:= Order{
		id: id,
		amount: amount,
		status: status,	
		createdAt: createdAt,
	}
	return &myOrder
}

func main() {
	fmt.Println("A struct is a collection of fields/Different DataTypes.")
	// var order Order = Order{
	// 	id:      "1",
	// 	amount:  20.64,
	// 	status:  "Recieved",
	// 	// createdAt: time.Now(),
	// }
	// order.createdAt = time.Now()
	// order.changeStatus("Delivered!")
	// fmt.Println(order)


	// myOrder:=newOrder("1", 20.54, "recieved", time.Now())
	// fmt.Println(myOrder)


	// language:=struct{
	// 	name string
	// 	isGood bool 
	// }{"golang", true}

	// fmt.Println(language)

	myOrder:=Order{
		id:"1",
		amount: 123.434,
		createdAt: time.Now(),
		status: "Failed",
		customer: customer{name:"Suraj", phoneNum:1234567890},
	}
	fmt.Println(myOrder)
}

// IF we dont set any value to a field in the struct, the default value will be zero Value

// we should use the pointers to pass the variable not its copy
// HP@DESKTOP-9HGE6HE MINGW64 ~/OneDrive/Desktop/golang (main)
// $ go run 16_struct/main.go 
// A struct is a collection of fields/Different DataTypes.
// {1 20.64 Recieved {13991914956923999264 504201 0x7ff71438ff80}}

// HP@DESKTOP-9HGE6HE MINGW64 ~/OneDrive/Desktop/golang (main)
// $ go run 16_struct/main.go 
// A struct is a collection of fields/Different DataTypes.
// {1 20.64 Delivered! {13991915011061555464 1 0x7ff69724ff80}}