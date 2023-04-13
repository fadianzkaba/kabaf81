package main

import (
	"fmt"
	"math/rand"
)

func main() {
	a := func() {
		fmt.Println("Hello world")

	}
	a()
	fmt.Println("%T", a)

	fmt.Println(rand.Intn(6))

}
