package main

import "github.com/tsawler/myniceprogram/helpers"

func main() {
	s := []string{"Fadi", "Farah", "Francis", "Fabiana"}
	s = append(s, "Mum", "Dad")
	helpers.DisplayOutPut(s)
}
