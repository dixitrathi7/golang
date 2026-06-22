package main

import "fmt"

func main() {
	var username string = "Alice"
	fmt.Println(username)
	fmt.Printf("Hello %s Nice to meet you \nAnd type of %T\n", username, username)

	var isLoggedIn bool = true
	fmt.Printf("Is user %s is Logged in (%t)\n", username, isLoggedIn)

	var num1 uint8 = 255 // uint8 can hold values from 0 to 255, int8 can hold values from -128 to 127
	fmt.Printf("Value of num1 is %d\n", num1)

	var num2 float32 = 3.14345325234 // float64 can hold more precision than float32, but float32 is sufficient for most cases
	fmt.Printf("Value of num2 is %f\n", num2)

	var defaultvalues uint // default value of int is 0 for uninitialized variablesv = 0
	fmt.Printf("Default value of int is %d\n", defaultvalues)

	var defaultfloat float32 // default value of float32 is 0.00000 for uninitialized variables
	fmt.Printf("Default value of float32 is %f\n", defaultfloat)

	var defaultvaluestring string // default value of string is empty string ""
	fmt.Printf("Default value of string is %q \n", defaultvaluestring)

	var website = "www.google.com" // type inference, the type of website is inferred to be string
	fmt.Printf("Value of website is %s\n", website)

	number := 42 // := is a shorthand for declaring and initializing a variable, the type of number is inferred to be int
	fmt.Printf("Value of number is %d\n", number)

}
