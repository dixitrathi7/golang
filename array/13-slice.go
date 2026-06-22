package main

import "fmt"

func uniqueslice(arr []string) ([]string, map[string]int) {
	slice := make([]string, 0, len(arr))
	seen := make(map[string]struct{})
	frequency := make(map[string]int)

	for _, v := range arr {
		if _, exists := seen[v]; !exists {
			seen[v] = struct{}{}
			slice = append(slice, v)
		}
		frequency[v]++
	}
	return slice, frequency
}

func main() {

	arr := []string{"MG", "AP", "OR", "BN", "MG", "AP", "OR", "BN", "MG", "AP", "OR", "BN", "MG", "AP", "OR", "BN"}
	fmt.Println(arr)

	slice, frequency := uniqueslice(arr)

	fmt.Println("Frequency of each element:")
	for k, v := range frequency {
		fmt.Printf("%s: %d\n", k, v)
	}

	fmt.Println("Unique elements in the slice:", slice)
	// for _, v := range slice {
	// 	fmt.Printf("%s\n", v)
	// }

	// 	for i, r := range arr {
	// 		fmt.Printf("%d: %s\n", i, r)
	// 	}

}
