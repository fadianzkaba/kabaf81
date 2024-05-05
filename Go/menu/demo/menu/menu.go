package menu

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var in = bufio.NewReader(os.Stdin)

type menuItem struct {
	name   string
	prices map[string]float64
}

type menu []menuItem

func (m menu) print() {
	for _, item := range m {
		fmt.Println(item.name)
		fmt.Println(strings.Repeat("-", 10))
		for size, price := range item.prices {
			fmt.Printf("%10s%10.2f\n", size, price)
		}
	}
}

func (m *menu) addItem() {
	fmt.Println("Please enter the items that you want to add to the list")
	name, _ := in.ReadString('\n')
	*m = append(*m, menuItem{name: name, prices: make(map[string]float64)})
}

// AddItem Add item to the Menu
func AddItem() {
	data.addItem()

}

// PrintMenu Display the data
func PrintMenu() {
	data.print()
}
