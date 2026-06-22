// Store numbers into a slice:
// When user enters: 0

package main

import "fmt"

func main() {

	var numbers []int
	for {
		var num int
		fmt.Print("Enter a number into slice (0 to stop): ")
		fmt.Scan(&num)
		if num == 0 {
			break
		}
		numbers = append(numbers, num)
	}
	fmt.Println("Numbers in slice:", numbers)
}
