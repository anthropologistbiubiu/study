package main

import "fmt"

func main() {

	fmt.Println(nthUglyNumber(10))
}

func nthUglyNumber(n int) int {

	var dp = make([]int, 0)
	dp[0] = 1
	var next int
	var p2, p3, p5 = 0, 0, 0
	var mp = map[int]struct{}{}
	var count int = 1
	for count <= n {
		mp[dp[i-1]] = struct{}{}
		dp[i], next = min(p2, p3, p5, dp, mp)
		if next == p2 {
			p2++
		} else if next == p3 {
			p3++
		} else if next == p5 {
			p5++
		}
		count++
	}
	fmt.Println(dp)
	return dp[n-1]
}

func min(p2, p3, p5 int, dp []int) (int, int) {
	ans := dp[p2] * 2
	next := p2
	if ans > dp[p3]*3 {
		ans = dp[p3] * 3
		next = p2
	} else {

	}
	return ans, next
}
