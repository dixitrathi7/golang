package main

import "fmt"

func canVote(age int) bool {
	if age >= 18 {
		return true
	} else if age < 18 {
		return false
	} else {
		return false
	}

}

// func eligablityfoAL(age int) bool {
// 	if age >= 21 {
// 		return true
// 	} else {
// 		return false
// 	}
// }

func main() {
	age := 45

	if canVote(age) == true {
		fmt.Println("You can Vote")
	} else {
		fmt.Println("You cannot Vote")
	}

}
