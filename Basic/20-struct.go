package main

import "fmt"

func main() {

	dog := struct {
		breed string
		age   int
	}{
		breed: "Labrador",
		age:   5,
	}
	fmt.Println(dog)
	fmt.Println(dog.breed)
	fmt.Println(dog.age)
}
