// find out even number

package main

import "fmt"

func main() {

	var numbers = []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	var evenNumbers []int

	for _, number := range numbers {
		if number%2 == 0 {
			evenNumbers = append(evenNumbers, number)
		}
	}

	fmt.Println("Even numbers:", evenNumbers)
}
