package main

import "fmt"

func main() {

	maps := map[string]string{
		"MG": "mango",
		"AP": "apple",
		"OR": "orange",
		"BN": "banana",
	}
	fmt.Println(maps)
	fmt.Println(len(maps))

	fmt.Println("update value the key AP")
	maps["AP"] = "avocado"

	delete(maps, "BN")

	fmt.Println(maps)

	fmt.Println("Print the length of the map")
	fmt.Println(len(maps))
}
