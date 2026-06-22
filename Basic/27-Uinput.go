package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Print("Enter your name: ")
	scanner.Scan()

	name := scanner.Text()
	fmt.Printf("Hello, %s!", name)
}
