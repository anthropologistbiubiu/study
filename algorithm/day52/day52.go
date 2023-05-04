package main

import "fmt"

//字符流中第一个不重复的字符

var s string
var set map[byte]int = make(map[byte]int)

func Insert(ch byte) {

	s += string(ch)
	set[ch]++
}

func FirstAppearingOnce() byte {

	for _, v := range []byte(s) {
		if set[v] == 1 {
			return v
		}
	}
	return '#'
}
func main() {
	var s = "abc孙"
	for _, v := range s {
		fmt.Printf("%+v  %T\n", v, v)
	}
	for _, v := range []byte(s) {
		fmt.Printf("%+v  %T\n", v, v)
	}
}
