package main

import "fmt"

func main() {

	type test struct {
		fname  string
		sname  string
		number int
	}

	t1 := test{
		fname:  "Fadi",
		sname:  "Kaba",
		number: 1,
	}

	fmt.Print(t1)

}
