package main

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
}
