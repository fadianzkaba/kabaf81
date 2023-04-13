package main

import (
	"errors"
	"fmt"
	"net/http"
)

const portNumber = ":9991"

func Home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello to the Home Page")
}

func About(w http.ResponseWriter, r *http.Request) {
	sum := addValues(2, 2)

	_, _ = fmt.Fprintf(w, fmt.Sprintf("This is About us Page %d", sum))
}

// Divide function is about Dividing the input
func Divide(w http.ResponseWriter, r *http.Request) {

	f, err := divideValues(100.0, 10.0)
	if err != nil {
		fmt.Fprintf(w, "Cannot divide by 0")
	}

	fmt.Fprintf(w, fmt.Sprintf("%f divided by %f is %f", 100, 0, f))

}

func divideValues(x, y float32) (float32, error) {
	if y <= 0 {
		err := errors.New("Cann't divide by 0")
		return 0, err
	}
	result := x / y
	return result, nil
}

func addValues(x, y int) int {
	return x + y
}

func main() {

	http.HandleFunc("/", Home)
	http.HandleFunc("/About", About)
	http.HandleFunc("/Divide", Divide)

	fmt.Println(fmt.Sprintf("Starting the appliation on port %s", portNumber))
	_ = http.ListenAndServe(portNumber, nil)

}
