package main

import "testing"

func TestDivide(t *testing.T) {
	_, err := divide(10.00, 0)
	if err != nil {
		t.Error("Got an error when should not have")
	}
}
