package main

import (
	"bufio"
	"demo/menu"
	"fmt"
	"os"
	"strings"
)

var in = bufio.NewReader(os.Stdin)

func main() {
loop:
	for {
		fmt.Println("1) Print Menu")
		fmt.Println("2) Add item")
		fmt.Println("3 Exit")

		choice, _ := in.ReadString('\n')

		switch strings.TrimSpace(choice) {
		case "1":
			menu.PrintMenu()

		case "2":
			menu.AddItem()
		case "3":
			break loop
		default:
			fmt.Println("\nPlease choice a valid option\n")
		}

	}
}
