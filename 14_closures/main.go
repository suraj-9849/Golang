package main

import "fmt"

func counter() func() int {
	count := 0

	// Anonymous Functions
	return func() int {
		count += 1
		return count
	}
}

func main() {
	increment := counter()
	fmt.Println(increment())
	fmt.Println(increment())
}