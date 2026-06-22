package main

import "fmt"

func main() {

	maps := make(map[string]int)
	{
		maps["MG"] = 100
		maps["AP"] = 200
		maps["OR"] = 300
		maps["BN"] = 400
	}

	fmt.Println(maps)

	fmt.Println(maps)

	maps["Ap"] = 500

	fmt.Println(maps)

	delete(maps, "BN")

	fmt.Println(maps)

}
