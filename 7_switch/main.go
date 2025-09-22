package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Switch Case:")
	// in golang there is no requirement of writing the break keyword golang handles it automatically!
	i := 2
	switch i {
	case 1:
		{
			fmt.Println("One")
			break
		}
	case 2:
		{
			fmt.Println("Two")
			break
		}
	default:
		{
			fmt.Println("Can be other than 1 or 2")
		}
	}

	// Multiple Condition Switch Case:
	switch time.Now().Weekday() {
	case time.Sunday, time.Saturday:
		{
			fmt.Println("Lets Enjoy!")
		}
	case time.Monday, time.Tuesday, time.Thursday, time.Wednesday:
		{
			fmt.Println("Boring Days")
		}
	case time.Friday:
		{
			fmt.Println("Waiting For Week Day")
		}
	}

	// type Switch:
	// any Type in go: interface{}
	whoAmI := func(i interface{}) {
		switch i := i.(type) {
		case int:
			{
				fmt.Printf("i is Integer %v\n", i)
			}
		case string:
			{
				fmt.Printf("i is String %v\n", i)
			}
		case bool:
			{
				fmt.Printf("i is Boolean %v\n", i)
			}
		case float32:
			{
				fmt.Printf("i is Float32 %v\n", i)
			}
		default:
			{
				fmt.Println("I dont Know the Type")
			}
		}
	}
	whoAmI(4)
}
