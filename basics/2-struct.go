package main

import (
	"log"
	"time"
)

type UserStu struct {
	FirstName   string
	LastName    string
	PhoneNumber string
	Age         int
	BirthDate   time.Time
}

// Variable starts with Caps is available for other packages
var String string

func struc() {
	user := UserStu{
		FirstName:   "Rakesh",
		LastName:    "K",
		PhoneNumber: "1111111",
	}
	log.Println(user.FirstName, user.LastName)
}
