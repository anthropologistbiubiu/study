package main

import "fmt"

// 给你一个整数数组 nums ，除某个元素仅出现 一次 外，其余每个元素都恰出现 三次 。请你找出并返回那个只出现了一次的元素。
// 你必须设计并实现线性时间复杂度的算法且使用常数级空间来解决此问题
func singleNumber(nums []int) int {

	var dp = make([]int, 32)
	for _, v := range nums {
		for j := 0; j < 32; j++ {
			dp[j] += (v >> j) & 1
		}
	}
	var result int
	fmt.Println(dp)
	for i, _ := range dp {
		fmt.Println(i)
		result |= (dp[i] % 3 << i)
	}
	return result
}

// 示例 1：
//输入：nums = [2,2,3,2]
//输出：3

//示例 2：
//输入：nums = [0,1,0,1,0,1,99]
//输出：99

func main() {
	//fmt.Println(singleNumber([]int{2, 2, 3, 2}))
	//fmt.Println(singleNumber([]int{0, 1, 0, 1, 0, 1, 99})) //输入：nums = [0,1,0,1,0,1,99]
	fmt.Println(singleNumber([]int{-2, -2, 1, 1, 4, 1, 4, 4, -4, -2}))
	// 对负数的处理是错误的
	/*
		var nums = []int{0, 1, 0, 1, 0}
		var tem int = 1
		var ans int
		for i := len(nums) - 1; i >= 0; i-- {
			if i == len(nums)-1 {
				tem *= 1
			} else {
				tem *= 2
			}
			if nums[i]&1 == 1 {
				ans += tem
			}
		}
	*/
	//var dp = []int{-2, -2, 1, 1, 4, 1, 4, 4, -4, -2}
}
