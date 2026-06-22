// for loop is the only loop in Go. It has three components separated by semicolons:
// 1. Initialization: executed before the loop starts (e.g., i := 0)
// 2. Condition: evaluated before each iteration; if false, the loop ends (e.g., i < 5)
// 3. Post statement: executed after each iteration (e.g., i++)
package main

import "fmt"

func loops(value int) int {
	for value < 8 {
		fmt.Println("Current value:", value)
		value++
	}
	return value
}

func main() {
	//i := 0
	//
	//for i <= 4 {
	//	fmt.Println("Current value of i:", i)
	//	i++
	//}

	//for i := 0; i < 4; i++ {
	//	fmt.Println("Current value of i:", i)
	//}

	fmt.Println("Return value from loops function:", loops(1))

}
