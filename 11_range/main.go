package main

import "fmt"

func main() {
	fmt.Println("Range")
	// Range: It is used for iterating over the DS
	nums := []int{1, 2, 3, 4}
	i := 0
	for i < len(nums) {
		fmt.Println(nums[i])
		i += 1
	}
	sum := 0
	for idx, num := range nums {
		fmt.Println(num, idx)
		sum += num
	}
	fmt.Println("sum: ", sum)
	m:=map[string]string{"price":"12dollars", "age":"12yrs"}
	fmt.Println("Key\t value")
	 for k,v:=range m {
		fmt.Println(k, "\t", v)
	 }

	 s:="suraj"
	 fmt.Println("Index\tCharacter")
	 // idx: starting byte of rune
	 // ch: unicode point rune
	 for idx, ch:=range s{
		fmt.Println(idx,"\t",string(ch))
	 }
}
