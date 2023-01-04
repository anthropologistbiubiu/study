package main

import (
	"fmt"
	"math"
)

func Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
func minCost() {
	var n int
	fmt.Scan(&n)
	var arr = make([]int, n)
	var dp = make([]int, n)
	for k := 0; k < n; k++ {
		fmt.Scan(&arr[k])
		dp[k] = 1
	}
	maxlen := dp[0]
	for i := 1; i < n; i++ {
		for j := 0; j < i; j++ {
			if arr[i] > arr[j] {
				dp[i] = Max(dp[j]+1, dp[i])
			}
		}
		maxlen = Max(maxlen, dp[i])
	}
	fmt.Println("maxlen", maxlen)
}

// 最小花费爬楼梯
func minCostClimbingStairs(cost []int) int {
	// write code here
	var length int = len(cost)
	var dp []int = make([]int, length+1)
	dp[0], dp[1] = 0, 0
	for i := 2; i <= length; i++ {
		dp[i] = int(math.Min(float64(dp[i-1]+cost[i-1]), float64(dp[i-2]+cost[i-2])))
	}
	return dp[length]
}
func main() {
	var cost = []int{10, 15, 20}
	var cost1 = []int{1, 100, 1, 1, 1, 100, 1, 1, 100, 1}
	res := minCostClimbingStairs(cost)
	fmt.Println(res)
	result := minCostClimbingStairs(cost1)
	fmt.Println(result)
}
