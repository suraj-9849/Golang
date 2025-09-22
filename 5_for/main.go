package main

import "fmt"
// for: is the Only construct in the GO->for Looping (no while, no do-while)
func main() {
	
	// fmt.Println("For!");
	// while loop:
	// i:=1;
	// for i<=3 {
	// 	fmt.Println(i)
	// 	i = i+1
	// }

	// Classic For Loop:
	// for i :=0; i<=3;i++ {
	// 	if i==2{
	// 		continue;
	// 	}
	// 	fmt.Printf("Hello world: %v\n", i)
	// }


	// in new go version : we got the range Keyword 
	for i:=range 3{
		fmt.Printf("%v\n",i)
	}
}