package main

import (
	"fmt"
	"log"
)

func main() {
	var color string
	color = "Green"

	log.Println("This is the First", color)
	changeColor(&color)
	log.Println("This is the Second", color)

}

func changeColor(s *string) {
	log.Println("This is the address of", s)
	newValue := "Red"
	log.Println()
	fmt.Println()

	*s = newValue
}
