package main

import "fmt"

const (
	C1 = iota + 1
	C2
	_
	C3
)

func main() {
	fmt.Println(C1, C2, C3)
}
