package main

import "fmt"

type Student struct {
	Name  string
	Age   int
	Class int
}

func updateValue(s *Student) string {
	s.Name = "Bob"
	return s.Name
}

func main() {
	s := Student{Name: "Alice", Age: 20, Class: 10}
	fmt.Println(&s) // get the address of the pointer variable s

	//	updatedname := updateValue(&s)
	s.Name = "Bob"
	updatedname := s.Name

	fmt.Println("Here is the Updated Value of Student:", updatedname)

	fmt.Println("Here is the original Value of Student:", s.Name)

}
