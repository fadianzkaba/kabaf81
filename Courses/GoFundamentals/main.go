package main

import (
	"fmt"
)

type Foo struct {
	bar string
  }

  func main() {
	list := []Foo{{"A"}, {"B"}, {"C"}}

	cp := make([]*Foo, len(list))
	for i, value := range list {
	  cp[i] = &value
	  fmt.Println(i, &value)
	}

    for i :=0; i< 3; i++{
		fmt.Println(i, cp, cp[i])
	}



  }