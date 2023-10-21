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
	var result int32
	for i, _ := range dp {
		if dp[i]%3 != 0 {
			result |= (int32(dp[i] % 3 << i))
		}
	}
	return int(result) // 解决一下负数的问题
}

func singleNumber1(nums []int) int {
	//var dp = make([]int, 32)
	var result int32
	for i := 0; i < 32; i++ {
		var total int
		for _, num := range nums {
			total += (num >> i) & 1
		}
		if total%3 != 0 {
			result |= (1 << i)
		}
	}
	return int(result)
}

func singleNumber_2(nums []int) int {
	ans := int32(0)
	for i := 0; i < 32; i++ {
		var total int
		for _, num := range nums {
			total += (num) >> i & 1
		}
		if total%3 > 0 {
			ans |= 1 << i
		}
	}
	// 64位的操作系统 int是2的64-1
	//-4本身表示为 二进制表示为  11111111111111111111111111111100
	// int 类型可以存储这个二进制为正数  4294967292
	// 因此需要使用int32,得到对应的负数
	return int(ans)
	/*
		具体的 int32 的范围如下：
		最小值：-2,147,483,648（-2^31）
		最大值：2147483647（2^31 - 1）
	*/
}

func main() {
	//fmt.Println(singleNumber([]int{-2, -2, 1, 1, 4, 1, 4, 4, -4, -2}))
	//fmt.Println(singleNumber1([]int{-2, -2, 1, 1, 4, 1, 4, 4, -4, -2}))
	fmt.Println(singleNumber_2([]int{-2, -2, 1, 1, 4, 1, 4, 4, -4, -2}))
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
