package main

import "fmt"

func main() {
	i:=3
	if i==2{
		fmt.Println("I am at 2 bro")
	} else if i == 3 {
        fmt.Println("I am at 3 bro")
    } else {
        fmt.Println("I am not at 2 bro")
    } 
	fmt.Println("Hello world")



	// New thing:
	if age:=14; age>=18{
		fmt.Println("I am adult boy")
	}else {
		fmt.Println("I am young boy")
	}

	// go doesn't have the Ternary Operator we should use the above methods only!
}