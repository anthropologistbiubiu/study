package main

import "fmt"

func main() {

	fmt.Println(nthUglyNumber(3))
}

//给你一个整数 n ，请你找出并返回第 n 个 丑数 。

//说明：丑数是只包含质因数 2、3 和/或 5 的正整数；1 是丑数。

func nthUglyNumber(n int) int {

	var dp = make([]int, n)
	dp[0] = 1
	var next int
	var p2, p3, p5 = 0, 0, 0
	for i := 1; i < n; i++ {
		dp[i], next = min(p2, p3, p5, dp)
		if next == p2 {
			p2++
		} else if next == p3 {
			p3++
		} else if next == p5 {
			p5++
		}
	}
	return dp[n-1]
}

func min(p2, p3, p5 int, dp []int) (int, int) {
	m := dp[p2] * 2
	next := p2
	if m > dp[p3]*3 {
		m = dp[p3] * 3
		next = p3
	}
	if m > dp[p5]*5 {
		m = dp[p5] * 5
		next = p5
	}
	return m, next
}
