package main

import "fmt"

func main() {
	secondinminute := 60
	minuteinhour := 60
	secondinhour := secondinminute * minuteinhour

	fmt.Printf("Number of second in one hour is:= %d\n", secondinhour)
	fmt.Println("add two number %d and %d is %f\n", secondinminute, minuteinhour, float64(secondinminute+minuteinhour))
}

// outputs
// Number of second in one hour is:= 3600
// add two number %d and %d is %f
// 60 60 120
