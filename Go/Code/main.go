package main

import (
	"github.com/pluralsight/webservice/models"
	"fmt"
)

func main() { 
	var u models.User

	u.ID = 1
	u.FirstName = "Fadi"
	u.SecondName = "Kaba"

	fmt.Println(u)

}
