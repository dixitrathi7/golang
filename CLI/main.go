package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	// If the user provides arguments, handle them as a one-shot command.
	if len(os.Args) > 1 {
		switch strings.ToLower(os.Args[1]) {
		case "hello":
			name := "World"
			if len(os.Args) > 2 {
				name = strings.Join(os.Args[2:], " ")
			}
			printGreeting(name)
		case "age":
			if len(os.Args) < 3 {
				fmt.Println("Usage: cmd-tool age <number>")
				return
			}
			printAge(os.Args[2])
		case "help", "--help", "-h":
			printHelp()
		case "repl":
			runREPL()
		default:
			fmt.Printf("Unknown command: %s\n\n", os.Args[1])
			printHelp()
		}
		return
	}

	// No arguments: start the interactive shell.
	runREPL()
}
