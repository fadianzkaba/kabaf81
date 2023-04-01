package main

import "log"

func main() {
	var mystring string
	mystring = "Green"

	log.Println("myString is set to", mystring)
	changeUsingPointer(&mystring)
	log.Println("myString is set to", mystring)

}

func changeUsingPointer(s *string) {
	log.Println(s)

	newValue := "Red"
	*s = newValue

}
