package main

import "fmt"

func main() {
	dividend, divisor := 10, 5
	fmt.Printf("%v Dividend by %v is %v\n", dividend, divisor, divide(dividend, divisor))

	dividend, divisor = 10, 0
	fmt.Printf("%v Dividend by %v is %v\n", dividend, divisor, divide(dividend, divisor))

}

func divide(dividend, divisor int) int {
	defer func() {
		if msg := recover(); msg != nil {
			fmt.Println(msg)
		}

	}()
	return dividend / divisor
}
