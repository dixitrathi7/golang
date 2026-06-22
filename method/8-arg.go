package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {

	if len(os.Args) <= 2 {
		fmt.Println("pass at least 2 integer command line arguments")
	} else {
		// Convert string arguments to integers
		num1, _ := strconv.Atoi(os.Args[1])
		num2, _ := strconv.Atoi(os.Args[2])
		sum := num1 + num2

		fmt.Println("The sum of the two command line integer arguments is : ", os.Args[1], "+", os.Args[2], "=", sum)
	}
}
