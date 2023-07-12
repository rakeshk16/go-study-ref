package main

import (
	"log"
	"sort"
)

func Slice() {
	var mySliceString []string
	mySliceString = append(mySliceString, "Trevor")
	mySliceString = append(mySliceString, "John")

	log.Println(mySliceString)

	var mySliceInt []int
	mySliceInt = append(mySliceInt, 1)
	mySliceInt = append(mySliceInt, 33)
	mySliceInt = append(mySliceInt, 3)
	log.Println(mySliceInt)

	sort.Ints(mySliceInt)

	log.Println(mySliceInt)

	number := []int{1, 4, 5, 6, 7, 3, 34, 54, 65}
	log.Println(number)
	log.Println(number[0:2])

	string1 := []string{"1", "4", "5", "6", "7", "3", "34", "54", "65"}
	log.Println(string1)
	log.Println(string1[0:2])
}
