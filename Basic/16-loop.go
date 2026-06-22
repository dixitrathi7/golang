package main

import "fmt"

func main() {
	i := 0
	for {
		fmt.Println("Current index: %d", i)
		i++
		if i >= 10 {
			break
		}
	}
}
