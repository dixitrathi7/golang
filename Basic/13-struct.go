package main

import "fmt"

type Employee struct {
	Name     string
	Age      int
	Title    string
	IsRemote bool
}

func main() {
	employee1 := Employee{
		Name:     "Alice",
		Age:      30,
		Title:    "DevOps Engineer",
		IsRemote: true,
	}

	fmt.Println("Employee Struct:", employee1)
	//fmt.Println("Employee Name:", employee1.Name)
	//fmt.Println("Employee Age:", employee1.Age)
	//fmt.Println("Employee Title:", employee1.Title)
	//fmt.Println("Is Employee Remote?", employee1.IsRemote)
}
