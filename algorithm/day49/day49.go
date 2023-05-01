package main

import (
	"fmt"
)

/*
地上有一个m行n列的方格，从坐标 [0,0] 到坐标 [m-1,n-1] 。
一个机器人从坐标 [0, 0] 的格子开始移动，它每次可以向左、右、上、下移动一格（不能移动到方格外），
也不能进入行坐标和列坐标的数位之和大于k的格子。
例如，当k为18时，机器人能够进入方格 [35, 37] ，因为3+5+3+7=18。
但它不能进入方格 [35, 38]，因为3+5+3+8=19。请问该机器人能够到达多少个格子？
*/

func movingCount1(threshold int, rows int, cols int) int {

	var m int = rows
	var n int = cols
	var k int = threshold
	var set = make(map[[2]int]struct{}, 0)
	var count int
	var tail int = -1
	var stack = make([][2]int, m+n+2)
	set[[2]int{0, 0}] = struct{}{}
	var x, y int
	for x < m && Help(x, y) <= k {
		tail++
		stack[tail] = [2]int{x, y}
		set[[2]int{x, y}] = struct{}{}
		x++
		count++
	}
	for tail >= 0 {
		x = stack[tail][0]
		y = stack[tail][1]
		tail--
		_, ok := set[[2]int{x + 1, y}]
		if Help(x+1, y) <= k && x+1 < m && !ok {
			set[[2]int{x + 1, y}] = struct{}{}
			tail++
			stack[tail] = [2]int{x + 1, y}
			count++
		}
		_, ok = set[[2]int{x, y + 1}]
		if Help(x, y+1) <= k && y+1 < n && !ok {
			set[[2]int{x, y + 1}] = struct{}{}
			tail++
			stack[tail] = [2]int{x, y + 1}
			count++
		}

	}

	return count
	// write code here
}
func Help(x, y int) int {

	var res int
	for x != 0 {
		res += x % 10
		x /= 10
	}
	for y != 0 {
		res += y % 10
		y /= 10
	}
	return res
}
func movingCount(m int, n int, k int) int {

	var set = make(map[[2]int]struct{}, 0)
	var count int
	var tail int = -1
	var stack = make([][2]int, m+n+2)
	var x, y int
	for x < m && x+y <= k {
		tail++
		stack[tail] = [2]int{x, y}
		set[[2]int{x, y}] = struct{}{}
		x++
		count++
	}
	for tail >= 0 {
		x = stack[tail][0]
		y = stack[tail][1]
		tail--
		_, ok := set[[2]int{x + 1, y}]
		if x+1+y <= k && x+1 < m && !ok {
			set[[2]int{x + 1, y}] = struct{}{}
			tail++
			stack[tail] = [2]int{x + 1, y}
			count++
		}
		_, ok = set[[2]int{x, y + 1}]
		if x+y+1 <= k && y+1 < n && !ok {
			set[[2]int{x, y + 1}] = struct{}{}
			tail++
			stack[tail] = [2]int{x, y + 1}
			count++
		}
	}

	return count
}
func main() {

	/*
			输入：m = 2, n = 3, k = 1
		输出：3
		示例 2：

		输入：m = 3, n = 1, k = 0
		输出：1
	*/
	fmt.Println(movingCount1(15, 20, 20))
	fmt.Println(movingCount(20, 20, 15))
}
