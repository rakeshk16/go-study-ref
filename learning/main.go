package main

import (
	"log"
)

func main() {
	isTrue := false
	num := 998
	if !isTrue && num < 90 {
		log.Println("isTrue is true")
	} else if num == 999 {
		log.Println("Number", num)
	} else if num != 999 || !isTrue {
		log.Println("Number", num)
	} else {
		log.Println("isTrue is false")
	}

	switch num {
	case 998:
		log.Println(num)
	case 9989:
		log.Println(num)
	default:
		log.Println("Default")
	}
}
