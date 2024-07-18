package main

import "fmt"

func main() {

	//EX1 (Print Function)
	fmt.Println("Hey, I am beginning with Golang!")

	//EX2 (Variable)
	var smsSendingLimit int
	var costPerSMS float64
	var hasPermission bool
	var usernamee string
	fmt.Printf("%v %.2f %v %q\n", smsSendingLimit, costPerSMS, hasPermission, usernamee)

	//EX3 (Short Variable Declaration)
	messageStart := "Happy birthday! You are now"
	age := 21
	messageEnd := "years old!"

	fmt.Println(messageStart, age, messageEnd)

	//EX4 (Same line Declaration)
	averageOpenRate, displayMessage := .23, "is the average open rate of your messages"

	fmt.Println(averageOpenRate, displayMessage)

	//EX5 (Type Sizes)
	accountAgeFloat := 2.6
	accountAgeInt := int(accountAgeFloat)

	fmt.Println("Your account has existed for", accountAgeInt, "years")

	//EX6 (Concatenation)
	var username string = "presidentSkroob"
	var password string = "12345"

	fmt.Println("Authorization: Basic", username+":"+password)

	//EX7 (Constants)
	const premiumPlanName = "Premium Plan"
	const basicPlanName = "Basic Plan"

	fmt.Println("plan:", premiumPlanName)
	fmt.Println("plan:", basicPlanName)

	//EX 8 (Computed Constants)
	const secondsInMinute = 60
	const minutesInHour = 60
	const secondsInHour = secondsInMinute * minutesInHour

	fmt.Println("number of seconds in an hour:", secondsInHour)

	//EX9 (Formatting Strings in Go)
	const name = "Saul Goodman"
	const openRate = 30.5

	msg := fmt.Sprintf("Hi %s, your open rate is %.1f percent\n", name, openRate)

	fmt.Print(msg)
}
