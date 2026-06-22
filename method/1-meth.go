package main

import "fmt"

type Student struct {
	Name string
	age  int
}

func (s Student) Getname(newname string) string {
	s.Name = newname
	return s.Name
}

func main() {

	s1 := Student{Name: "Alice", age: 20}

	name := s1.Getname("Bob")
	fmt.Println(name)
	fmt.Println(s1.Name)

}
