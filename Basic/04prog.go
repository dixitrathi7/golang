package main

import "fmt"

func main() {
	age := 20
	exage := 20.78

	aproxage := float64(age)

	stexage := int(exage)

	fmt.Printf("Type of above variable is \n%f\n%d", aproxage, stexage)
}
