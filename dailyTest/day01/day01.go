package main

import (
	"fmt"
	"reflect"
)

type S struct {
	Name    string
	Age     int
	Address *int
}

func main() {
	a := S{
		Name:    "aa",
		Age:     1,
		Address: new(int),
	}
	b := S{
		Name:    "aa",
		Age:     1,
		Address: new(int),
	}

	fmt.Println(a == b)
	fmt.Println(reflect.DeepEqual(a, b))
}
