package main

import (
	"fmt"
	"math/rand"
)

type action func(s score) (result score, turnInOver bool)

type score struct {
	player, opponent, thisTurn int
}

func main() {

	a := score(2)
	a2 := action(getMeNumber)
	fmt.Println(a2(a))

	d := getMeNumber
	fmt.Println(d(a))

	c := action(getMeNumber)
	fmt.Println(c(a))

}

func getMeNumber(c score) (score, bool) {
	return c * 2, true
}

func getAnotherNumber(c score) score {
	return c * 8
}

func roll(s score) (score, bool) {
	outcome := rand.Intn(6) + 1 // A random int in [1, 6]
	if outcome == 1 {
		return score{s.opponent, s.player, 0}, true
	}
	return score{s.player, s.opponent, outcome + s.thisTurn}, false
}
