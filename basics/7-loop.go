package main

import (
	"log"
)

func Loop() {
	for i := 0; i <= 5; i++ {
		log.Println(i)
	}

	animals := []string{"dog", "cat", "lion", "tiget"}

	// _ ignores the value here it is index
	for _, animal := range animals {
		log.Println(animal)
	}

	animals1 :=
		make(map[string]string)
	animals1["Dog"] = "ggg"
	animals1["Cat"] = "cat"
	for _, animal1 := range animals1 {
		log.Println(animal1)

	}

	var firstName = "Once upon a time in the night"
	for i, l := range firstName {
		// String is slice of bytes
		log.Println(i, ":", l)
	}

	type Userloop struct {
		FirstName string
		LastName  string
		Age       int
	}

	var users []Userloop

	users = append(users, Userloop{"Rakesh", "K", 27})
	users = append(users, Userloop{"Rahul", "K", 28})
	users = append(users, Userloop{"jj", "K", 30})
	users = append(users, Userloop{"hh", "K", 25})
	for _, l := range users {
		log.Println(l)
	}
}
