package main

import "fmt"

type Person struct {
	Name string
}

func main() {
	a := Person{Name: "Raj"}
	b := &a          // b is a pointer to a, so it points to the same underlying data as a
	b.Name = "Rohan" // is this update the data addressed by b, which is the same as the data addressed by a?
	// answer is yes, because b is a pointer to a, so it points to the same underlying data as a,
	// so when we update the data addressed by b, we are also updating the data addressed by a,
	// because they are the same data.

	fmt.Println(a.Name) // still "Raj"
	fmt.Println(b.Name)
}
