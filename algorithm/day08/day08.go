package main

import "fmt"

//给定一个整数数组 nums ，找到一个具有最大和的连续子数组（子数组最少包含一个元素），返回其最大和。
//
//示例:
//
//输入: [-2,1,-3,4,-1,2,1,-5,4]
//输出: 6
//解释: 连续子数组 [4,-1,2,1] 的和最大，为 6。
//进阶
//
//如果你已经实现复杂度为 O(n) 的解法，尝试使用更为精妙的分治法求解。
<<<<<<< HEAD

func maxSubArray(nums ...int) int {
	//res := make([]int, len(nums))
	maxValue := 0
	//from, to := 0, 0
	//for i := 0; i < len(nums); i++ {
	//	j = i
	//	sum := 0
	//	for j < len(nums) {
	//		if sum+nums[j] > sum {
	//			sum += nums[j]
	//		} else {
	//			if sum > maxValue {
	//				maxValue = sum
	//				from = i
	//				to = j
	//			}
	//			break
	//		}
	//		j++
	//	}
	//}
	dp := make([]int, len(nums))
	i := 1
	dp[0] = nums[0]
	for i < len(nums) {
		if dp[i-1]+nums[i] > nums[i] {
			dp[i] = dp[i-1] + nums[i]
		} else {
			dp[i] = nums[i]
		}
		if dp[i] > maxValue {
			maxValue = dp[i]
		}
		i++
	}
	//fmt.Println(from, to)
	return maxValue
}
func main() {

	nums := []int{-2, 1, -3, 4, -1, 2, 1, -5, 4}
	fmt.Println(maxSubArray(nums...))

}
=======
>>>>>>> 3ffb4361f2d9a132248437aa8fc76cf36d0b9bfb
