package main

import "fmt"

/*
给你一个整数数组 nums ，找到其中最长严格递增子序列的长度。

子序列 是由数组派生而来的序列，删除（或不删除）数组中的元素而不改变其余元素的顺序。例如，[3,6,2,7] 是数组 [0,3,1,6,2,2,7] 的子序列。

输入：nums = [10,9,2,5,3,7,101,18]
输出：4
解释：最长递增子序列是 [2,3,7,101]，因此长度为 4 。

输入：nums = [7,7,7,7,7,7,7]
输出：1
*/
func lengthOfLIS(nums []int) int {

	var dp = make([]int, len(nums))
	dp[0] = 1
	for i := 1; i < len(nums); i++ {
		if nums[i] > nums[i-1] {
			dp[i] = dp[i-1] + 1
		} else {
			dp[i] = dp[i-1]
		}
	}
	return dp[len(nums)-1]
}

func main() {
	//var nums = []int{0, 1, 0, 3, 2, 3}
	var nums = []int{7, 7, 7, 7, 7, 7}
	fmt.Println(lengthOfLIS(nums))
}
