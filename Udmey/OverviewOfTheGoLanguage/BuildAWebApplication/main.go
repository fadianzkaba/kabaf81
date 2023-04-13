package main

import (
	"fmt"
	"runtime"
)

func main() {

	a, b, c, d := runtime.Caller(0)

	fmt.Println("A:", a, "\nB:", b, "\nC:", c, "\nD:", d)

	/*http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		n, err := fmt.Fprintf(w, "Hellow, world")
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println(fmt.Sprintf("Number of bytes: %d", n))
	})

	_ = http.ListenAndServe(":9991", nil)*/
}
