package main

import "fmt"

func main() {
	var nums []int
	fmt.Println(nums)

	var size int
	fmt.Print("Enter the size of the array: ")
	fmt.Scanln(&size)

	nums = make([]int, size)

	fmt.Println(nums)

	fmt.Println("Enter the array numbers:")

	for i := 0; i < len(nums); i++ {
		fmt.Printf("Enter number %d: ", i+1)
		fmt.Scanln(&nums[i])
	}
	fmt.Println("Print Array:", nums)

	fmt.Println("Array Length:", len(nums))

	for i := 0; i < len(nums); {
		fmt.Println("looping through the array: ", nums[i])
		i++
	}

}
