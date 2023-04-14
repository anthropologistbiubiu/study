package main

import (
	"fmt"
	"reflect"
)

type People struct {
	Age int64
}

func main() {
	p := &People{
		Age: 30,
	}

	fmt.Println(reflect.ValueOf(p).Kind()) // ouput: ptr
	value := reflect.Indirect(reflect.ValueOf(p))
	fmt.Println(value.Kind()) // output: struct
}
