package main

import "fmt"

func main() {

	maps := map[string]int{
		"MG": 100,
		"AP": 200,
		"OR": 300,
		"BN": 400,
	}

	if val, exists := maps["AP"]; exists {
		fmt.Printf("The value of the key AP is %d\n", val)
	} else {
		fmt.Println("The key AP does not exist in the map")
	}
}
