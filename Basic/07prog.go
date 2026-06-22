package main

import "fmt"

func main() {

	number := 2

	if length := number; length > 5 {
		fmt.Println("number is too long")
	} else if length < 5 {
		fmt.Println("number is too short")
	} else {
		fmt.Println("number is just right")
	} // length is only available in this block only
}
