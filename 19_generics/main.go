package main

import "fmt"

// func printSlice(items []int) {
// 	for _, item:=range items {
// 		fmt.Println(item)
// 	}
// }

// func printSlice2(items []string) {
// 	
// }

func printSliceGeneric[T comparable] (items []T){
	for _, item:=range items {
		fmt.Println(item)
	}
}

type Stack[T string | int] struct{
	elements []T
}



func main() {
// This is not the better way of using by creating 2 functions which does the same activity: so we use the Generics
	// printSlice([]int{1, 2, 3, 4, 5})
	// printSlice2([]string{"1", "2", "3", "4", "5"})

	printSliceGeneric([]int{1,2,3,4,5})
	printSliceGeneric([]string{"golang","js","ts"})


	newStack := Stack[int]{
		elements: []int{1, 2, 3, 4, 8},
	}
	newStack2 := Stack[string]{
		elements: []string{"bmw","bugatti"},
	}
	fmt.Println(newStack)
	fmt.Println(newStack2)
}