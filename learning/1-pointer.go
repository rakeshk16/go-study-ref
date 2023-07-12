package main

import (
	"log"
)

func pointer() {
	var myString string
	myString = "Green"
	log.Println(myString)
	changeUsingPointer(&myString)
	log.Println(myString)
}

func changeUsingPointer(s *string) {
	newValue := "Red"
	*s = newValue
}
