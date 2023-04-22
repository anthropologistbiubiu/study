package main

import (
	"fmt"
)

/*
给定一个字符串，验证它是否是回文串，只考虑字母和数字字符，可以忽略字母的大小写。

说明：本题中，我们将空字符串定义为有效的回文串。

示例 1:

输入: "A man, a plan, a canal: Panama"
输出: true
示例 2:

输入: "race a car"
输出: false
*/

func isValid(s string) string {
	var res string
	for tail, _ := range s {
		if s[tail] >= 65 && s[tail] <= 90 || s[tail] >= 97 && s[tail] <= 122 || s[tail] >= 48 && s[tail] <= 57 {
			res += string(s[tail])
		}

	}
	return res
}
func isPalindrome(s string) bool {

	s = isValid(s)
	fmt.Println(s)
	if len(s) == 0 {
		return true
	}
	var top, tail = 0, len(s) - 1
	for top <= tail {
		fmt.Println(s[top])
		fmt.Println(s[tail])
		if s[top] >= 48 && s[top] <= 57 {
			if s[top] != s[tail] {
				return false
			}
		}
		if !(s[top] == s[tail] || s[top] == s[tail]+32 || s[top]+32 == s[tail]) {
			return false
		}
		tail--
		top++
	}
	return true
}
func main() {

	//var s = "fghj , 8"
	//fmt.Println(string(s[2]))
	//fmt.Printf("%T\n", (s[2]))
	//fmt.Println(string(104))
	//fmt.Println(len(s))
	var assist = "0P"
	fmt.Println(isPalindrome(assist))
}
