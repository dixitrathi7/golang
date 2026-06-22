package main

import "fmt"

type Student struct {
	name  string
	grade []int
	age   int
}

func (s *Student) getage(age int) {
	s.age = age
	//	return s.age
}

func main() {

	s1 := Student{"Daksh", []int{90, 80, 85}, 25}
	fmt.Println(s1)
	s1.getage(24)
	fmt.Println("Now the age through getage method is: ", s1)

}
