package main

import (
	"log"
	"time"
)

type User struct {
	FirstName   string
	LastName    string
	PhoneNumber string
	Age         int
	DOB         time.Time
}

func (m *User) callingsomething() (string, string) {
	return m.FirstName, m.LastName
}

func main() {

	myVar1 := User{
		FirstName: "Fadi",
		LastName:  "Kaba",
	}

	myVar2 := User{
		FirstName: "Farah",
		LastName:  "Kaba",
	}

	log.Println(myVar1.callingsomething())
	log.Println(myVar2.callingsomething())

}
