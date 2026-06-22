package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter your name: ")
	name, _ := reader.ReadString('\n')

	name = strings.TrimSpace(name)

	fmt.Print("Enter your age: ")
	var age int
	fmt.Scanf("%d", &age)

	fmt.Printf("Hello, %s", name)
	fmt.Printf("Your age is %d", age)
}
