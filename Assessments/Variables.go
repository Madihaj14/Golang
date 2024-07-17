package main

import "fmt"

func main() {

	//EX1 (Print Function)
	fmt.Println("Hey, I am beginning with Golang!")

	//EX2 (Variable)
	var smsSendingLimit int
	var costPerSMS float64
	var hasPermission bool
	var username string
	fmt.Printf("%v %.2f %v %q\n", smsSendingLimit, costPerSMS, hasPermission, username)

	//EX3 (Short Variable Declaration)
	messageStart := "Happy birthday! You are now"
	age := 21
	messageEnd := "years old!"

	fmt.Println(messageStart, age, messageEnd)

	//EX4 (Same line Declaration)
	averageOpenRate, displayMessage := .23, "is the average open rate of your messages"

	fmt.Println(averageOpenRate, displayMessage)
}
