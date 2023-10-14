package main

import "fmt"

func coinChange(coins []int, amount int) int {

	var dp = make([]int, amount+1)
	dp[0] = 0
	for i := 1; i <= amount; i++ {
		var min int
		for _, v := range coins {
			if i >= v && dp[i-v] > 0 || i-v == 0 {
				if min >= dp[i-v]+1 || min == 0 {
					min = dp[i-v] + 1
				}
			}
		}
		dp[i] = min
	}
	fmt.Println(dp)
	if amount > 0 && dp[amount] == 0 {
		return -1
	}
	return dp[amount]
}

func main() {
	// 示例 1：
	//
	//输入：coins = [1, 2, 5], amount = 11
	//输出：3
	//解释：11 = 5 + 5 + 1
	//fmt.Println(coinChange([]int{1, 2, 5}, 11))
	//示例 2：
	//
	//fmt.Println(coinChange([]int{2}, 3))
	//输入：coins = [2], amount = 3
	//输出：-1
	//示例 3：
	//
	fmt.Println(coinChange([]int{2, 5, 10, 1}, 27))
	//输入：coins = [1], amount = 0
	//输出：0
	// 我应该尝试写一个单元测试函数，用来提高代码的质量
	//[2,5,10,1]
}
