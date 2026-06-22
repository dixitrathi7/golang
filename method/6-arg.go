package main

import (
	"fmt"
	"os"
)

func main() {

	fmt.Println("The command line arguments are: ", os.Args[0:])

	i := 0
	for i < len(os.Args) {
		fmt.Println("Argument number: ", i, " is ", os.Args[i])
		i++
	}
}
