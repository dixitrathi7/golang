package main

import "fmt"

func main() {

	type person struct {
		name     string
		age      int
		location string
	}

	p1 := person{
		name:     "John",
		age:      30,
		location: "New York",
	}

	fmt.Println(p1)
	p2 := person{"Jane", 28, "Chicago"}
	fmt.Println(p2)
	fmt.Println(person{name: "Alice", age: 25, location: "Los Angeles"})
	fmt.Println(person{"Bob", 35, "Miami"})

	fmt.Println(p1.age)
	p1.age = 20
	fmt.Println(p1.age)

}
