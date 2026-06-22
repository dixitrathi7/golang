package main

import "fmt"

func main() {
	username := "Alice"
	password := "password123"
	fmt.Println("Username:password", username+":"+password)

	amount, message := 100.50, "only february payment"

	fmt.Println("Here is the amount and message:", amount, message)

}
