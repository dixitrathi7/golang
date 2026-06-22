package main

import (
	"fmt"
)

type Student struct {
	Name  string
	age   int
	grade []int
	class string
}

func (s Student) AverageGrade() float64 {
	sum := 0
	for _, g := range s.grade {
		sum += g
	}
	return float64(sum) / float64(len(s.grade))
}

func main() {
	s1 := Student{
		Name:  "Jack",
		age:   20,
		grade: []int{85, 90, 78},
		class: "Computer Science",
	}
	fmt.Println("Average of student's grade is:", s1.AverageGrade())
}
