// You can edit this code!
// Click here and start typing.
package main

import "fmt"

func main() {
	type user struct {
		ID    int
		fname string
		sname string
	}
	u1 := user{
		ID:    1,
		fname: "Francis",
		sname: "Kaba",
	}
	u2 := user{
		ID:    2,
		fname: "Fabiana",
		sname: "Kaba",
	}
	if u1.ID == u2.ID {
		fmt.Println("Same User")
	} else if u1.fname != u2.fname {
		fmt.Println("They are not the same")
	} else {
		fmt.Println("They are not the same")
	}

}