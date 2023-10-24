package main

import "fmt"

func main() {

	fmt.Println(nthUglyNumber(11))
}

func nthUglyNumber(n int) int {

	var dp = make([]int, n)
	dp[0] = 1
	var p2, p3, p5 = 0, 0, 0
	for i := 1; i < n; i++ {
		ans, next := min(p2, p3, p5, dp)
		dp[i] = ans
		if next == p2 {
			p2++
		} else if next == p3 {
			p3++
		} else if next == p5 {
			p5++
		}
		if ans == dp[i-1] {
			i--
		}
	}
	return dp[n-1]
}

func min(p2, p3, p5 int, dp []int) (int, int) {
	fmt.Println("dp", dp)
	ans := dp[p2] * 2
	next := p2
	if ans > dp[p3]*3 {
		ans = dp[p3] * 3
		next = p2
	}
	if ans > dp[p5]*p5 {
		ans = dp[p5] * p5
		next = p5
	}
	return ans, next
}
