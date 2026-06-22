package main

import "fmt"

func main() {

	slice := []int{4, 5, 6, 7, 8, 9}
	fmt.Println(slice)

	s := slice[1:3]
	fmt.Println(s)

	slice = append(slice, 10, 11, 12)
	fmt.Println("Here we append 10, 11, and 12 into the slice s,")
	fmt.Println(s)
	fmt.Println(slice)

}
