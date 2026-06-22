package main

import (
	"fmt"
	"os"
)

func main() {

	if len(os.Args) == 1 {
		fmt.Println("No command line arguments were passed")
	} else {
		i := 1
		for i < len(os.Args) {
			fmt.Println("hello : ", os.Args[i], "Sir Good Day")
			i++
		}
	}
}
