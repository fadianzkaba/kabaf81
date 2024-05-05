package main

import "fmt"

type myInt struct {
	id int
	name string
}



func main() {


	my1 := myInt
	ans := mi.isEven()

	fmt.Println(ans)

}

func (i myInt) isEven() bool {
	return int(i)%2 == 0
}
