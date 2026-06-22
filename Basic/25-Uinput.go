package main

import "fmt"

func main() {
	var name string
	var age int
	fmt.Printf("Enter you name: ")
	fmt.Scanln(&name)

	fmt.Printf("Enter your age: ")
	fmt.Scan(&age)

	fmt.Println("Hello, ", name)
	fmt.Println("Your age is ", age)

}
