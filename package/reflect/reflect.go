package main

import (
	"fmt"
	"reflect"
)

type People struct {
	name string
	Age  int64
}

func main() {
	p := &People{
		name: "sunweiming",
		Age:  30,
	}

	fmt.Println(reflect.ValueOf(p)) // ouput: ptr
	value := reflect.Indirect(reflect.ValueOf(p)).Type()
	fmt.Println(value)
	for i := 0; i < value.NumField(); i++ {
		p := value.Field(i)
		fmt.Println(p)
	}
}
