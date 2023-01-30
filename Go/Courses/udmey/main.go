package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type person struct {
	FirstName  string `json:"firstName"`
	SecondName string `json:"secondName"`
	HairColor  string `json:"HairColor"`
	HasDog     bool   `json:"HasDog"`
}

func main() {
	myJson := `
	[
		{
			"firstName":"FName1",
			"secondName":"SName1",
			"HairColor":"Black",
			"HasDog":true
		},
		{
			"firstName":"FName2",
			"secondName":"SName2",
			"HairColor":"Blue",
			"HasDog":false
		}
	]`

	var unmarshalled []person

	err := json.Unmarshal([]byte(myJson), &unmarshalled)

	if err != nil {
		log.Println("Error unmarshelling json", err)
	}

	log.Printf("unmarshalled: %v", unmarshalled)

	var mySlice []person

	var m1 person
	m1.FirstName = "M1FName"
	m1.SecondName = "M1SName"
	m1.HairColor = "M1Black"
	m1.HasDog = true

	mySlice = append(mySlice, m1)

	var m2 person
	m2.FirstName = "M2FName"
	m2.SecondName = "M2SName"
	m2.HairColor = "M2Black"
	m2.HasDog = false

	mySlice = append(mySlice, m2)

	NewJson, err := json.MarshalIndent(mySlice, " ", "       ")

	if err != nil {
		fmt.Printf("Error")
	}

	log.Printf(string(NewJson))

}
