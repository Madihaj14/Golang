//EX1 Conditionals

package main

import "fmt"

func main() {
	messageLen := 10
	maxMessageLen := 20
	fmt.Println("Trying to send a message of length:", messageLen, "and a max length of:", maxMessageLen)

	// don't touch above this line

	if messageLen <= maxMessageLen {
		fmt.Println("Message sent")
	} else {
		fmt.Println("Message not sent")
	}

	//EX2 shorter way of writing
	//foreg, instead of writing:
	length := getLength(email)
	if length < 1 {
		fmt.Println("Email is invalid")
	}
	//you can write:
	if length := getLength(email); length < 1 {
		fmt.Println("Email is invalid")
	}

}
