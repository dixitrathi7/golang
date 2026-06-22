package main

import "fmt"

func main() {
	height := 4

	if height > 8 {
		fmt.Println("You are tall")
	} else if height >= 6 {
		fmt.Println("You are just right")
	} else if height >= 4 {
		fmt.Println("You are short")
	} else {
		fmt.Println("Something is wrong with your height")
	}
}
