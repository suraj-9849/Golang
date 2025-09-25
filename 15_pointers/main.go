package main

import "fmt"

func main() {
	fmt.Println("Pointers allow you to modify a variable's value by passing its memory address (reference) to a function.")
	// num := 1
	// changeNum(num)
	// fmt.Println("After the ChangeNUm the num in main: ", num)
	
	num:=1
	changeNumPointer(&num)
	fmt.Println("After the ChangeNUm the num in main: ", num)

}


// by value will be passed a copy
// func changeNum(num int) {
// 	num = 5
// fmt.Println("In change Num",num)
// }

func changeNumPointer(num *int){
	*num=5
	fmt.Println("In change Num",*num)
}

// $ go run 15_pointers/main.go 
// In change Num 5
// After the ChangeNUm the num in main:  1


// go run 15_pointers/main.go 
// Pointers allow you to modify a variable's value by passing its memory address (reference) to a function.
// In change Num 5
// After the ChangeNUm the num in main:  5

