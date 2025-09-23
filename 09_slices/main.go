package main

import (
	"fmt"
	"slices"
)

func main() {
	// Slices are used when we dont have idea about the size (Dynamic Size)
	// Mostly used in Golang
	// we get +useful methods
	arr := []int{}
	var arr2 []int
	fmt.Println(arr2)
	fmt.Println(arr)
	// uninitalized slice is called nil
	fmt.Println(arr == nil)  //false
	fmt.Println(arr2 == nil) // true
	var nums = make([]int, 2)
	fmt.Println(len(nums))
	fmt.Println(nums == nil) // false

	// 	[]int{} → empty non-nil slice (len=0, cap=0, but has a valid slice header).

	// var s []int → nil slice (len=0, cap=0, no backing array, header is nil).

	// Both print as [], but they are not the same.

	// and in the make we have 3 arguments where: 1 slice , 2 len, 3 capacity
	// capacity -> Maximum elements that can fit before reallocation

	var nums2 = make([]int, 2, 10)
	fmt.Println(cap(nums2))

	// it the capacity increases: Go automatically allocates a new, larger array (usually doubling the size).
	// 	So after your append:

	// Old backing array: cap=10 (discarded when exceeded).

	// New backing array: larger (at least 25, probably 32).

	// Slice points to the new array.
	nums2 = append(nums2, 1, 2, 3, 4, 4, 7, 9, 5, 3, 2, 4, 6, 7, 4, 2, 2, 5, 6, 7, 5, 3, 3, 2)
	fmt.Println(cap(nums2))
	fmt.Println(nums2)

	// copy the slice:
	var newSlice  = make([]int, 0,10)
	newSlice = append(newSlice, 2)
	var newSlice2  = make([]int, len(newSlice	))

	fmt.Println(newSlice, newSlice2)

	// copy(dest, src)
	copy(newSlice2, newSlice)
	fmt.Println(newSlice2)


	// Slice Operator:
	 fmt.Println(nums2[2:6])
	 fmt.Println(nums2[1:])
	 fmt.Println(nums2[:8])

	 //Slice Package: Slices
	 var nums1 = []int{1,2}
	 var nums3 = []int{1,2}
	 fmt.Println(slices.Equal(nums1, nums3))

	 // 2D slices:
	 var twod = [][]int{{1,2,3},{4,5,6}}
	 fmt.Println(twod)
}
