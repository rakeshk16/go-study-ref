package main

import "log"

type UserMap struct {
	FirstName string
	LastName  string
}

func (u *UserMap) displayValuesOfUserMap() string {
	return u.FirstName
}

func Map() {
	myMap := make(map[string]string)
	myMap1 := make(map[string]int)
	myStructMap := make((map[string]UserMap))

	myMap["Dog"] = "Doggie"

	myMap["Dog2"] = "Doggie@"

	myMap["Dog"] = "Doggie7"

	log.Println(myMap["Dog"])

	myMap1["Dog"] = 1

	myMap1["Dog2"] = 12

	log.Println(myMap1["Dog"])

	myStructMap["1"] = UserMap{
		FirstName: "Rakesh",
		LastName:  "K",
	}

	log.Println(myStructMap["1"].FirstName)

}
