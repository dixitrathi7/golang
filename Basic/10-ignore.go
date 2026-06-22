package main

import "fmt"

func ignorelastname() (string, string) {
	return "Dixit", "Rathore"
}

func ignoreint(num1, num2 int) (int, int) {

	numsum := num1 + num2

	nummul := num1 / num2

	return numsum, nummul
}

func main() {

	num1 := 40
	num2 := 30

	firstname, lastname := ignorelastname()

	onlyfirstname, _ := ignorelastname()

	numsum, _ := ignoreint(num1, num2)

	fmt.Println("Only firstname is : ", onlyfirstname)

	fmt.Println("Full name of candidate is : ", firstname, lastname)

	fmt.Println("sum of two number is : ", numsum)

}
