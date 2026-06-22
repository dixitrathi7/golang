package main

import "fmt"

type User struct {
	Name string
	age  int
}

func createuser(name string, age int) User {
	return User{Name: name, age: age}
}

func main() {

	user := createuser("Dixit", 24)
	fmt.Println(user.Name)
	fmt.Println(user.age)
}s
