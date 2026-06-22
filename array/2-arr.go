package main

import "fmt"

func main() {

	var arr [3]int

	arr[0] = 1
	arr[1] = 2
	arr[2] = 3
	fmt.Println(arr)

	arr[1] = 20

	fmt.Println(arr)

	fmt.Println(arr[2])

	fmt.Println(len(arr))
}
