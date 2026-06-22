package main

import "fmt"

func main() {

	var day string

	fmt.Println("Enter a day of the week:")
	fmt.Scanln(&day)

	switch day {

	case "Monday":
		fmt.Println("Start of the work week.")
	case "Tuesday":
		fmt.Println("Second day of the work week.")
	case "Wednesday":
		fmt.Println("Midweek day.")
	case "Thursday":
		fmt.Println("Almost the weekend.")
	case "Friday":
		fmt.Println("Last day of the work week.")
	case "Saturday", "Sunday":
		fmt.Println("It's the weekend!")
	}
}
