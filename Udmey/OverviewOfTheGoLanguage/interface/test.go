package main

import "fmt"

type abc struct {
	abc1 cdf
}

type cdf interface {
	Apply1() string
	fig() string
}

func Apply1() string {
	return "Hello"
}

func fig() string {
	return "Test"
}

func main() {

	a := abc{
		abc1: cdf.Apply1(),
	}

	fmt.Printf(a.abc1.Apply1()

}
