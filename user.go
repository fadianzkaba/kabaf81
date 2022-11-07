package model 

import "fmt"

type User struct {
	ID int
	fname string
	sname string 
}

var (
	users []*User
	nextID=1
)

func GetUsers() []*User {
	return users
}

func AddUser (u User) (User, error) {
	u.ID = nextID
	nextID++
	users = append (users, &u)
	return u, nil
}