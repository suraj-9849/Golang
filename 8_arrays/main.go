package main

import "fmt"

func main() {
	fmt.Println("Arrays!")

	// var nums []int;
	// i:=0;
	
    // fmt.Println("Enter integers (one per line). Press Ctrl+D (Unix) or Ctrl+Z (Windows) to finish:")

	// for{
	// 	var num int;
	// 	_, err:= fmt.Scanf("%v\n",&num);
	// 	if err!=nil{
	// 		break;
	// 	}
	// 	nums = append(nums, num)
	// }
	
    // fmt.Printf("Length of Array is : %v\n", len(nums))

	// for i<len(nums) {
	// 	fmt.Println(nums[i])
	// 	i+=1;
	// }

	arr:= [4]int{1,2,3,4}
	fmt.Println(arr)

	// 2D-Arrays:
	arr2D := [2][2]int{{1,2},{3,4}}
	fmt.Println(arr2D)

	// Arrays can be used when we know we have the fixed size, Memory Optimisation, Constant time access 
}