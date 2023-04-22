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
		a ^= b
		b = carry
	}
	return a

}
func getSum1(a int, b int) int {
	if a == 0 {
		return b
	} else if b == 0 {
		return a
	}
	carry := a & b
	if carry != 0 {
		for carry != 0 {
			a = a ^ b
			carry = a & b << 1
		}

	} else {
		return a ^ b
	}
	return a
}
func main() {
	fmt.Println(getSum(3, 2))
	fmt.Println(getSum(3, 2))
}
