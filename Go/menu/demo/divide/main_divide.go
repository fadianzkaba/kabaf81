package mainv1

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func mainv1() {

	for {
		fmt.Println("1) Print Menu")
		fmt.Println("2) Add item")
		fmt.Println("3 Exit")

		in := bufio.NewReader(os.Stdin)
		s, _ := in.ReadString('\n')
		s = strings.TrimSpace(s)
		s = strings.ToUpper(s)

		i, _ := strconv.Atoi(s)
		if i < 4 {
			i--
			displayMenu(i)

		} else {
			fmt.Println("\nThank you\n", strings.Repeat("!", 30))
			break
		}
	}

}

func displayMenu(i int) {
	type menuItem struct {
		name   string
		prices map[string]float64
	}

	menu := []menuItem{
		{name: "Coffee", prices: map[string]float64{"Large": 1.60, "Medium": 1.50, "Small": 1.40}},
		{name: "Tea", prices: map[string]float64{"Hot Tea": 1.50, "Milk Tea": 1.60, "Black Tea": 1.60}},
		{name: "Iced Coffee", prices: map[string]float64{"Large": 1.70, "Medium": 1.60, "Small": 1.5}},
	}

	item := menu[i]
	fmt.Println(item.name)
	fmt.Println(strings.Repeat("-", 10))
	for size, price := range item.prices {
		fmt.Printf("%10s%10.2f\n", size, price)

	}
}
