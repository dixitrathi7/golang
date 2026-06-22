package main

import "fmt"

func main() {

	var age int

	fmt.Println("Enter your age:")
	fmt.Scanln(&age)

	switch {
	case age <= 18:
		fmt.Println("You are a minor.")

	case age > 18 && age < 65:
		fmt.Println("You are an adult.")

	case age >= 65:
		fmt.Println("You are a senior.")

	default:
		fmt.Println("Invalid age entered.")
	}
}
