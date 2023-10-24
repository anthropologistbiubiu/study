package main

import "fmt"

func main() {

	fmt.Println(nthUglyNumber(10))
}

func nthUglyNumber(n int) int {

	var dp = make([]int, n)
	dp[0] = 1
	var next int
	var p2, p3, p5 = 0, 0, 0
	var mp = map[int]struct{}{}
	for i := 1; i < n; i++ {
		mp[dp[i-1]] = struct{}{}
		dp[i], next = min(p2, p3, p5, dp, mp)
		if next == p2 {
			p2++
		} else if next == p3 {
			p3++
		} else if next == p5 {
			p5++
		}
	}
	fmt.Println(dp)
	return dp[n-1]
}

func min(p2, p3, p5 int, dp []int, mp map[int]struct{}) (int, int) {
	fmt.Println("dp", dp, p2, p3, p5)
	m := dp[p2] * 2
	var result int
	next := p2
	var next2 int
	var next3 int
	if m > dp[p3]*3 {
		m = dp[p3] * 3
		next = p3
		next2 = p2
	} else {
		next2 = p3
	}
	if m > dp[p5]*5 {
		m = dp[p5] * 5
		next = p5
	} else {
		if dp[p5]*5 < dp[next2]*next2 {
			next3 = next2
			next2 = p5
		} else {
			next3 = p5
		}
	}
	for _, v := range []int{next, next2, next3} {
		if _, ok := mp[v]; ok {
			continue
		} else {
			result = v
		}
	}
	return dp[result] * result, result
}
