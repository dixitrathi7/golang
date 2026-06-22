package main

import "fmt"

type person struct {
	person_name string
	age         int
}

type Company struct {
	Company_Name string
	CEO          *person
}

func main() {
	p1 := person{person_name: "Alice", age: 30}
	// p2 := person{person_name: "Bob", age: 40}
	// p3 := person{person_name: "Charlie", age: 35}
	// p4 := person{person_name: "David", age: 28}

	c1 := Company{Company_Name: "Tech Corp", CEO: &p1}
	fmt.Println(p1.person_name)

	fmt.Println("the CEO name of the company is:", c1.CEO.person_name)

}
