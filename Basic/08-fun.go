package main

import "fmt"

func addstr(stra, strb string) string {
	return stra + strb
}

func addint(num1, num2 int) int {
	numsum := num1 + num2
	return numsum
}

func main() {
	fmt.Println("for concatenate the string: ", addstr("hello ", "world"))
	fmt.Println("for concatenate the string: ", addstr("go ", "Lang is fun"))

	fmt.Println("here is the sum of two numbers: ", addint(10, 20))
	fmt.Println("here is the sum of two numbers: ", addint(100, 200))
}
