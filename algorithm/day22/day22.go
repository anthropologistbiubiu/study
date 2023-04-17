package main

import "fmt"

/*
不使用运算符 + 和 - ​​​​​​​，计算两整数 ​​​​​​​a 、b ​​​​​​​之和。

示例 1:

输入: a = 1, b = 2
输出: 3
示例 2:

输入: a = -2, b = 3
输出: 1
*/

func getSum(a int, b int) int {

	if a == 0 {
		return b
	} else if b == 0 {
		return a
	}
	carry := 0
	for b != 0 {
		carry = a & b << 1
		a = a ^ b
		b = carry
	}

	return a
}
func main() {
	fmt.Println(getSum(100, 2))
}
