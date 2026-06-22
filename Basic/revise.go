package main

import "fmt"

type person struct {
	name string
	age  int
}

func main() {
	p1 := person{name: "Alice", age: 30}

	fmt.Println(p1.name)
}
