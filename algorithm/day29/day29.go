package main

import (
	"fmt"
	"math"
)

// 字符串之间的最短距离
/*
给定一个字符串 S 和一个字符 C。返回一个代表字符串 S 中每个字符到字符串 S 中的字符 C 的最短距离的数组。

示例 1:

输入: S = "loveleetcode", C = 'e'
输出: [3, 2, 1, 0, 1, 0, 0, 1, 2, 2, 1, 0]
说明:

- 字符串 S 的长度范围为 [1, 10000]。
- C 是一个单字符，且保证是字符串 S 里的字符。
- S 和 C 中的所有字母均为小写字母。
*/

func shortestToChar(s string, c rune) []int {

	asis := make([]int, 0)
	ret := make([]int, 0)
	for i, v := range s {
		if v == c {
			asis = append(asis, i)
		}
	}
	for i := 0; i < len(s); i++ {
		var length float64 = float64(len(s))
		for _, j := range asis {
			if math.Abs(float64(i-j)) < float64(length) {
				length = math.Abs(float64(i - j))
			}
		}
		ret = append(ret, int(length))
	}
	return ret
}

func main() {

	//输入: S = "loveleetcode", C = 'e'
	//输出: [3, 2, 1, 0, 1, 0, 0, 1, 2, 2, 1, 0]
	s := "loveleetcode"
	c := 'e'
	fmt.Printf("%+v\n", c)
	fmt.Printf("%T\n", c)
	fmt.Println(shortestToChar(s, c))
}
