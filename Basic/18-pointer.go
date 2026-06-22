package main

import "fmt"

func main() {
	hight := 7
	pointer := &hight

	fmt.Printf("The address of hight is: %d and the value of hight is %d \n", pointer, *pointer)
	*pointer = 10

	hight = 20

	secondhight := hight

	fmt.Println("Access value diretly :", hight)

	fmt.Println("Access value diretly :", secondhight)

	fmt.Println("Access value via pointer :", *pointer)

}
