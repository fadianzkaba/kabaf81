package main

import (
	"fmt"
)

func main() {

	// a:=18

	// if a==18 {
	// 	fmt.Println("You are 18 years old")

	// }else {
	// 	fmt.Println("You are not 18 yet")
	// }

	// if num:= rand.Intn(10); num>5{
	// 	fmt.Println("Number is greater than 5")
	// }else {
	// 	fmt.Println("Number is less 5")
	// }

	// b:=9

	// switch b {
	// case 5, 6, 9:
	// 	fmt.Println("Yes it is between 6.5.9")
	// case 10,20,30:
	// 	fmt.Println("Yes it is between 10,20,30")
	// 	fallthrough
	// case 70,80:
	// 	fmt.Println("Well the Number is not between 70,80 but it is: ", b)
	// }

	//For Loop

	numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	for i, num := range numbers {
		fmt.Println("Display the numbers", i, num)
	}

	// var b int

	fmt.Println(b)

	loop:
	b++
	if b < 20{
		fadi(b)
		goto loop
	}


}

func fadi(b int) {
	fmt.Println("Yes ", b)
}
