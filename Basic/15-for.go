package main

import "fmt"

func loops(value int) int {
	for value < 8 {
		value++
		if value == 4 {
			continue
		}
		fmt.Println("Current value:", value)
		//	value++
	}
	return value
}

func main() {

	fmt.Println("Return value from loops function:", loops(1))

	for i := range 5 {
		fmt.Println("Current index:", i)
	}
}
