package main

import "fmt"

func addFn(a int, b int) int{
	return a+b
}


func multi(a int, b int)(result int){
	result =a*b
	return 
}

func multiFn(a int, b int) (sub int, div int){
	sub = a-b
	div = a/b
	return 
}

// only one type is enough
func add2Fn(a,b int) int{
	return a+b
}

func getLanguages() (string, string, string){
	return "golang", "js", "rust"
}

// function getting function as the Parameter:
func processIt(fn func(a int))  {
	fn(42) 
}

func main() {
	fmt.Println("Functions")
	ans1:=addFn(4,5)
	fmt.Println("add:", ans1)
	ans2:=multi(4,5)
	fmt.Println("multiplication:", ans2)
	ans3, ans4:=multiFn(10,5)
	fmt.Printf("sub:%v\t div:%v\n", ans3, ans4)

	fn:=func (a int)  {
		fmt.Println(a)
	}
	processIt(fn)
}