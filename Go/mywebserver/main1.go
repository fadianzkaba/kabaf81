package main()

import "fmt"

type user struct { 
	ID int
	fname string 
	sname string
}

func main() { 

	u1 := user {
		ID = 1
		fname = "Fadi"
		sname = "Kaba"
	}

	u2 := user {
		ID = 2
		fname = "Farah"
		sname = "Wadee"
	}

	println (u1, u2)

}