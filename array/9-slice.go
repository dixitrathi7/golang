package main

import "fmt"

func main() {
	nums := []int{1, 2, 3, 4, 5}

	fmt.Println(nums) // [1 2 3 4 5]
	// remove the element of slice
	newslice := append(nums[:2], nums[3:]...)
	fmt.Println(newslice) // [1 2 4 5]
}
