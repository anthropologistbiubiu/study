package main

import "fmt"

/*
给定一个三角形 triangle ，找出自顶向下的最小路径和。

每一步只能移动到下一行中相邻的结点上。相邻的结点
在这里指的是 下标 与 上一层结点下标 相同或者等于 上一层结点下标 + 1 的两个结点。
也就是说，如果正位于当前行的下标 i ，那么下一步可以移动到下一行的下标 i 或 i + 1 。
*/
func minimumTotal(triangle [][]int) int {
	m := len(triangle)
	var result int
	var dp = make([][]int, m)
	for i := 0; i < m; i++ {
		dp[i] = make([]int, m)
	}
	dp[0][0] = triangle[0][0]
	for i := 1; i < m; i++ {
		dp[i][0] = dp[i-1][0] + triangle[i][0]
	}
	for j := 1; j < m; j++ {
		var i = j
		dp[i][j] = dp[i-1][j-1] + triangle[i][i]
	}
	fmt.Println("dp", dp)
	for i := 2; i < m; i++ {
		for j := 1; j < i; j++ {
			dp[i][j] = min(dp[i-1][j-1], dp[i-1][j]) + triangle[i][j]
		}
	}
	result = dp[m-1][0]
	for j := 0; j < m; j++ {
		if result > dp[m-1][j] {
			result = dp[m-1][j]
		}
	}
	return result
}
func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

// triangle = [[2],[3,4],[6,5,7],[4,1,8,3]]
// 输出：11

func main() {

	triangle := [][]int{{2}, {3, 4}, {6, 5, 7}, {4, 1, 8, 3}}
	fmt.Println(minimumTotal(triangle))
}
