package main

import (
	"errors"
	"fmt"
	"log"
)

func main() {

	result, err := divide(100, 0)

	if err != nil {
		log.Println(err)
		return
	}

	log.Println(result)
	fmt.Println(result)
}

func divide(x, y float32) (float32, error) {
	var result float32

	if y == 0 {
		return result, errors.New("Can't divide by 0")
	}

	result = x / y
	return result, nil
}
