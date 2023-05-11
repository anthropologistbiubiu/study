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

// 所以这里就是传值还是传引用的问题
// []string 传进去是错误的
func letterCombinations(digits string) []string {

	if len(digits) == 0 {
		return []string{}
	}
	index := 0
	var ans = []byte{}
	var result = []string{}
	traceback(digits, index, ans, result)
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

func traceback(digits string, index int, ans []byte, result []string) {
	if index == len(digits) {
		result = append(result, string(ans))
		fmt.Println(result)
		return
	}
	digit := string(digits[index])
	letters := Map[digit]
	for _, v := range []byte(letters) {
		ans = append(ans, v)
		traceback(digits, index+1, ans, result)
		ans = ans[:len(ans)-1]
	}
}

func Test(tem []string) {
	if len(tem) == 3 {
		return
	}
	tem = append(tem, "a")
	fmt.Println(tem)
	Test(tem)
}
func main() {
	var digits = "234"
	//输出：["ad","ae","af","bd","be","bf","cd","ce","cf"]
	// .   [ad    ae   af   bd   be   bf   cd   ce   cf]
	fmt.Println(letterCombinations(digits))
	//var tem = []string{}
	//Test(tem)
	//fmt.Println(tem)

}

// 今晚写完这道题
// 总结 由于go是传值的特点所以golang 的下层的全局变量是上层初始值的拷贝，所以在递归的函数中首先要考虑的是数据的传递类型。
// 可以简单比方为函数的入参数全局变量也就是栈的一层的局部变变量，函数外部的全局变量也就是整个递归过程中的全局变量。
