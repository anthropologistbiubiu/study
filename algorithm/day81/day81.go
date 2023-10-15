package main

import "fmt"

// dp[i] = nums[i]
func maxProduct(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	var mindp []int = make([]int, len(nums))
	var maxdp []int = make([]int, len(nums))
	var ans int
	for i := 0; i < len(nums); i++ {
		if i == 0 {
			mindp[i] = nums[i]
			maxdp[i] = nums[i]
		} else {
			var min, max int
			tem1 := mindp[i-1] * nums[i]
			tem2 := maxdp[i-1] * nums[i]
			if tem1 > tem2 {
				max = tem1
				min = tem2
			} else {
				max = tem2
				min = tem1
			}
			if max <= nums[i] {
				max = nums[i]
			}
			if min >= nums[i] {
				min = nums[i]
			}
			maxdp[i] = max
			mindp[i] = min
		}
		if ans < maxdp[i] || ans == 0 {
			ans = maxdp[i]
		}
	}
	return ans
}
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// 给你一个整数数组 nums ，请你找出数组中乘积最大的非空连续子数组（该子数组中至少包含一个数字），并返回该子数组所对应的乘积。
//
// 测试用例的答案是一个 32-位 整数。
//
// 子数组 是数组的连续子序列。

// 示例 1:
//输入: nums = [2,3,-2,4]
//输出: 6
//解释: 子数组 [2,3] 有最大乘积 6。
//示例 2:
//输入: nums = [-2,0,-1]
//输出: 0
//解释: 结果不能为 2, 因为 [-2,-1] 不是子数组。

func main() {
	fmt.Println(maxProduct([]int{-2, 3, 1, 0}))
}
