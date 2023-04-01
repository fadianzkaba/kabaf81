package main

import "log"

type User struct {
	Fname string
	Sname string
}

func (u *User) PrintValue() []string {
	u2 := User{
		Sname: "Kaba",
	}

	var u3 []string
	u3 = append(u3, u.Fname, u2.Sname)
	return u3
}

func main() {

	var user1 User
	user1.Fname = "John"
	user2 := User{
		Fname: "Fadi",
	}
	log.Println("Hello World", user1.PrintValue(), user2.PrintValue())
}
