package main

import "log"

func main() {
	myString := "Green"

	log.Println("MyString is set to", myString)
	changemycolor(&myString)
	log.Println("New Value", myString)
}

func changemycolor(s *string) {
	log.Println("Display Output", s)
	newvalue := "red"
	*s = newvalue
}
