package main

import "fmt"

type engine struct {
	cylinders  int
	horsepower int
	torque     int
	cc         int
}

type car struct {
	carname string
	model   int
	price   string
	color   string
	engine  engine
}

func main() {

	mycar := car{carname: "Mahindra Thar", model: 2022, price: "15Lakhs", color: "Red",
		engine: engine{ // engine is an anonymous struct field So you must: Repeat its structure Then provide values
			cylinders:  4,
			horsepower: 150,
			torque:     300,
			cc:         2000,
		},
	}
	fmt.Println(mycar)
}
