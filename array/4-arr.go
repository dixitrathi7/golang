package main

import "fmt"

func sumArr(arr [5]int) (int, float64) {
	sum := 0
	for i := 0; i < len(arr); i++ {
		sum = arr[i] + sum

	}
	avg := float64(sum) / float64(len(arr))
	return sum, avg
}

func getOcurrences(arr [20]int, num int) int {

	repetitions := 0
	for i := 0; i < len(arr); i++ {
		if arr[i] == num {
			repetitions++
		}
	}
	return repetitions
}

func main() {

	a := [5]int{5, 13, 65, 23, 44}

	o := [20]int{1, 1, 3, 1, 5, 3, 1, 8, 3, 1, 11, 1, 1, 4, 2, 1, 17, 8, 1, 20}

	fmt.Println(a)

	sum, avg := sumArr(a)

	fmt.Println("Sum of your array is:", sum)
	fmt.Println("Average of your array is:", avg)

	fmt.Println("Number of occurrences of 1 in the array is:", getOcurrences(o, 1))

}
