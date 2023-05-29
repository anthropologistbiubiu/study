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

	var dp = make([]int, 0)
	dp = append(dp, nums[0])
	var tail int = 0
	var maxlen int = 1
	fmt.Println("start", dp)
	for i := 1; i < len(nums); i++ {
		if nums[i] > dp[tail] {
			dp = append(dp, nums[i])
			tail++
			fmt.Println("dp1", dp, tail)
		} else {
			for tail >= 0 {
				if dp[tail] >= nums[i] {
					tail--
				} else {
					dp[tail+1] = nums[i]
					tail = len(dp) - 1
					//tail
					break
				}
			}
			if tail == -1 {
				tail++
				dp[tail] = nums[i]
				tail = len(dp) - 1

			}
			fmt.Println("dp2", dp, tail)
		}

		if len(dp) > maxlen {
			maxlen = len(dp)
		}
	}
	return maxlen
}

func main() {
	//var nums = []int{0, 1, 0, 3, 2, 3}
	//var nums = []int{7, 7, 7, 7, 7, 7}
	//var nums = []int{4, 10, 4, 2, 8, 9}
	//var nums = []int{10, 9, 2, 5, 3, 7, 101, 18}
	//var nums = []int{1, 3, 6, 7, 9, 4, 10, 5, 6}
	var nums = []int{18, 55, 66, 2, 3, 54}
	fmt.Println(lengthOfLIS(nums))
}
