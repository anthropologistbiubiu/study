package main

/*
电话号码的字母组合
给定一个仅包含数字 2-9 的字符串，返回所有它能表示的字母组合。答案可以按 任意顺序 返回。

给出数字到字母的映射如下（与电话按键相同）。注意 1 不对应任何字母。
链接：https://leetcode.cn/leetbook/read/top-interview-questions-medium/xv8ka1/

输入：digits = "23"
输出：["ad","ae","af","bd","be","bf","cd","ce","cf"]
*/
func letterCombinations(digits string) []string {

	return []string{}
}

/*
	const string letterMap[10] = {
	        "", // 0
	        "", // 1
	        "abc", // 2
	        "def", // 3
	        "ghi", // 4
	        "jkl", // 5
	        "mno", // 6
	        "pqrs", // 7
	        "tuv", // 8
	        "wxyz", // 9
	    };
*/

var Map = map[string]string{
	"0": " ",
	"1": " ",
	"2": "abc",
	"3": "defc",
	"4": "ghi",
	"5": "jkl",
	"6": "mno",
	"7": "pqrs",
	"8": "tuv",
	"9": "wxyz",
}

var result []string

func traceback(digits string, index int, ans string) {
	if index == len(digits) {
		result = append(result, ans)
	}
	digit := string(digits[index])
	letters := Map[digit]
	for i := 0; i < len(letters); i++ {

	}

}

func main() {

}

// 今晚写完这道题
