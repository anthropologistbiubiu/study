package main

import "fmt"

func main() {

	fmt.Println("hello world!")
	fmt.Println(int32(-4 >> 31))
	var a int = -4
	fmt.Println(a >> 31)
	var b int32 = -4
	fmt.Println(b >> 31)
}
