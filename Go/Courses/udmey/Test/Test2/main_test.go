package main

import "testing"

var tests = []struct {
	name     string
	dividend float32
	divisor  float32
	expected float32
	isErr    bool
}{
	{"valid-data", 100.0, 10.0, 10.0, false},
	{"invalid-data", 100.0, 0.0, 0.0, true},
	{"valid-data", 25.0, 5.0, 5.0, false},
}

func TestDivision(t *testing.T) {
	for _, tt := range tests {
		got, err := divide(tt.dividend, tt.divisor)
		if tt.isErr {
			if err == nil {
				t.Error("expected an error but didn't get one")
			}
		} else {
			if err != nil {
				t.Error("Expected an error but didn't get one", err.Error())
			}
		}

		if got != tt.expected {
			t.Errorf("Expected %f but got %f", tt.expected, got)
		}

	}
}

func TestDivide(t *testing.T) {

	_, err := divide(10.0, 1.0)

	if err != nil {
		t.Error("Got an error where I shouldn't")
	}

}

func TestBadDivide(t *testing.T) {

	_, err := divide(10.0, 0.0)

	if err != nil {
		t.Error("Dig not get an error when we should have")
	}

}
