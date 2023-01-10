package helpers

import "fmt"
type Users 

func DisplayOutPut(s []string) {

	for i := 0; i < len(s); i++ {
		fmt.Println("Name:", s[i])
	}
}
