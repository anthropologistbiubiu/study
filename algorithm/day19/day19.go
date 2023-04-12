package main

import (
	"fmt"
)

/*
给定一个整数，将其转化为7进制，并以字符串形式输出。

示例 1:

输入: 100
输出: "202"

示例 2:

输入: -7
输出: "-10"

注意: 输入范围是 [-1e7, 1e7] 。
*/
func convertToBase7(x int32) string {
	var ret []rune
	for x != 0 {
		temp := x % 7
		x = x / 7
		ret = append(ret, rune(temp))
	}
	res := ""
	for _, r := range ret {
		res += fmt.Sprintf("%c", r)
	}
	return res
}
func main() {
	fmt.Println(convertToBase7(100))

	fmt.Println(convertToBase7(-7))
}
