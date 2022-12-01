package main

import "fmt"

func main() {
	fmt.Println("Hello, World.")
	whatToSay := "Goodby, crul world"
	fmt.Println(whatToSay)
	i := 7
	fmt.Println(i)

	whatToSay1, else2 := saySomething()

	fmt.Println("The function returned", whatToSay1, else2)
}

func saySomething() (string, string) {
	return "Something", "else"
}
