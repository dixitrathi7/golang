package main

import "fmt"

func updateArr(arr *[5]int) {
	arr[3] = 300
}

func main() {

	a := [5]int{5, 13, 65, 23, 44}

	fmt.Println("Arr before update")
	fmt.Println(a)

	updateArr(&a)

	fmt.Println("Arr after update")
	fmt.Println(a)

	for i := len(a) - 1; i >= 0; i-- {
		fmt.Printf("value is this %d of index %d\n", a[i], i)
	}
}
