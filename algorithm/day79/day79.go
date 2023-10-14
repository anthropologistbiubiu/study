package day79

func coinChange(coins []int, amount int) int {

	var dp = make([]int, amount+1)
	dp[0] = 0
	for i := 1; i <= amount; i++ {
		for _, v := range coins {
		}
	}
	return 0
}
