package main

import (
	"log"
)

func main() {
	type User struct {
		FirstName string
		LastName  string
	}

	u := User{
		FirstName: "Fadi",
		LastName:  "Kaba",
	}

	v := []int{1, 2, 3, 4, 5, 6, 7}
	v = append(v, 4)

	log.Println(u.FirstName)
	log.Println(v)

	isTrue := true
	myNum := 100

	if myNum > 99 && !isTrue {
		log.Println("Is True")
	} else {
		log.Print("That didn't work")
	}

	switch myNum {
	case 98:
		log.Println("This is true")
	case 99:
		log.Println("This is NOT true")
	default:
		log.Println("This is default")

	}

}
