package main

import "fmt"

/*
电话号码的字母组合
给定一个仅包含数字 2-9 的字符串，返回所有它能表示的字母组合。答案可以按 任意顺序 返回。

给出数字到字母的映射如下（与电话按键相同）。注意 1 不对应任何字母。
链接：https://leetcode.cn/leetbook/read/top-interview-questions-medium/xv8ka1/

输入：digits = "23"
输出：["ad","ae","af","bd","be","bf","cd","ce","cf"]
*/
var result []string

func letterCombinations(digits string) []string {

	index := 0
	var ans = []byte{}
	traceback(digits, index, ans)
	return result
}

var Map = map[string]string{
	"0": " ",
	"1": " ",
	"2": "abc",
	"3": "def",
	"4": "ghi",
	"5": "jkl",
	"6": "mno",
	"7": "pqrs",
	"8": "tuv",
	"9": "wxyz",
}

func traceback(digits string, index int, ans []byte) {
	if index == len(digits) {
		result = append(result, string(ans))
		return
	}
	digit := string(digits[index])
	letters := Map[digit]
	for _, v := range []byte(letters) {
		ans = append(ans, v)
		traceback(digits, index+1, ans)
		ans = ans[:len(ans)-1]
	}

}

func main() {
	var digits = "23"
	//输出：["ad","ae","af","bd","be","bf","cd","ce","cf"]
	// .   [ad    ae   af   bd   be   bf   cd   ce   cf]
	letterCombinations(digits)
	fmt.Println(result)
}

// 今晚写完这道题
