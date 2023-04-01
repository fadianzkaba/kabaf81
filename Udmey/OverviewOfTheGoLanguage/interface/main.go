package main

import "fmt"

type Animal interface {
	say() string
	NumberOfLegs() int
}

type saySomething func() string

type Dog struct {
	Name  string
	Breed string
	woof  saySomething
}

func woof() string {
	return "Woof"
}

type Cat struct {
	Name      string
	color     string
	NoOfTeeth int
}

func main() {

	dog := Dog{
		Name:  "Bulldogs",
		Breed: "Blue",
		woof:  woof,
	}

	cat := Cat{
		Name:      "cat",
		color:     "Pink",
		NoOfTeeth: 5,
	}

	printInfo(&dog)
	printInfo(&cat)
	fmt.Println(dog.woof())

}

func printInfo(a Animal) {
	fmt.Println("This animal say", a.say(), "and have this nuumber of leg", a.NumberOfLegs())
}

func (d *Dog) say() string {
	return "Woof"
}
func (d *Dog) NumberOfLegs() int {
	return 4
}

func (d *Cat) say() string {
	return "Meo"
}
func (d *Cat) NumberOfLegs() int {
	return 4
}
