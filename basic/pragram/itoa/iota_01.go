package main

import "fmt"

const (
	C0 = iota
	C1 = iota
	C2 = iota
)

func main() {

	fmt.Println(C0, C1, C2) // "0 1 2"
}
