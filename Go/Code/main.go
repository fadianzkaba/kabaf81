package main

import (
	"fmt"
)

func main() {

	slice := make([]int, 4, 9)
	for i := 0; i < 7; i++ {
		slice = append(slice, i)
	}

	fmt.Println(slice, len(slice), cap(slice))

}
