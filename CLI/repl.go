package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func runREPL() {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("Welcome to cmd-tool interactive mode.")
	fmt.Println("Type 'help' to see available commands.")

	for {
		fmt.Print("jump$dixit > ")
		if !scanner.Scan() {
			break
		}

		line := strings.TrimSpace(scanner.Text())
		if line == "" {
			continue
		}

		parts := strings.Fields(line)
		switch strings.ToLower(parts[0]) {
		case "exit":
			fmt.Println("Goodbye!")
			return
		case "clear":
			fmt.Print("\033[H\033[2J")
		case "help":
			printHelp()
		case "hello":
			name := "World"
			if len(parts) > 1 {
				name = strings.Join(parts[1:], " ")
			}
			printGreeting(name)
		case "age":
			if len(parts) < 2 {
				fmt.Println("Usage: age <number>")
				continue
			}
			printAge(parts[1])
		default:
			fmt.Printf("Unknown command: %s\n", parts[0])
			fmt.Println("Type 'help' for a list of commands.")
		}
	}
}

func printGreeting(name string) {
	fmt.Printf("Hello, %s!\n", name)
}

func printAge(value string) {
	age, err := strconv.Atoi(value)
	if err != nil {
		fmt.Println("Please provide a valid number for age.")
		return
	}
	fmt.Printf("You are %d years old!\n", age)
}

func printHelp() {
	fmt.Println("cmd-tool usage:")
	fmt.Println("  cmd-tool           # start interactive mode")
	fmt.Println("  cmd-tool repl      # start interactive mode")
	fmt.Println("  cmd-tool hello NAME")
	fmt.Println("  cmd-tool age NUMBER")
	fmt.Println("  cmd-tool help")
	fmt.Println()
	fmt.Println("Interactive commands:")
	fmt.Println("  hello NAME  - Greet a person")
	fmt.Println("  age NUMBER  - Print your age")
	fmt.Println("  clear       - Clear the screen")
	fmt.Println("  help        - Show this help message")
	fmt.Println("  exit        - Exit interactive mode")
}

func conversionString(str string) []string {
	lower := strings.ToLower(str)
	words := strings.Fields(lower)
	return words
}
