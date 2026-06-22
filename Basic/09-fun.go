package main

import "fmt"

func avg(receivedmessage, totalusers int) int {
	receivedmessage = receivedmessage / totalusers
	return receivedmessage
}

func main() {

	receivedmessage := 400
	const totalusers = 50

	receivedmessage = avg(receivedmessage, totalusers)

	fmt.Println("Average of received messages : ", receivedmessage)
}
