package main

import "fmt"

func main () { 
	
	/*var f []int


	f = append(f, 1, 2, 3,4)
	type user struct {
		ID int
		fname string
		sname string 
	}

	var u user


	u.ID = 1
	u.fname = "Fadi"
	u.sname = "Kaba"

	fmt.Println(u.ID, u.fname, u.sname)
	fmt.Println("\n")
	result := cal (f)*/
	slice := map[string]int{"http":80, "https" : 443}

	for _, v := range slice {
		Sprintf(v)
	}
	
}

func cal (f []int) bool {
	fmt.Println (f)
	return true
}