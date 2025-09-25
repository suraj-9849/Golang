package main

import "fmt"

func sum(nums...int) int{
	total:=-0;
	for _, num := range nums {
		total += num
	}
	return total
}

func main() {
	fmt.Println("Variadic Functions")
	fmt.Println(1, 2, 3, 4, "Hello")
	result:=sum(1,2,3,4,6)
	fmt.Println(result)

	nums:=[]int{1,2,3,4}
	result2 :=sum(nums...) // spread operator same as the JS spread operator!
	fmt.Println(result2)
}