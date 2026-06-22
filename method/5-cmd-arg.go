package main

import (
	"fmt"
	"os"
)

func main() {

	fmt.Println("The command line arguments are: ", os.Args[0:])

	firstarg := os.Args[2]
	fmt.Println("The first command line argument is: ", firstarg)
}
