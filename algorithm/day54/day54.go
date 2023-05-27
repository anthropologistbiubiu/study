package main

import "fmt"

/*
给你一个整数数组 nums ，数组中的元素 互不相同 。返回该数组所有可能的子集（幂集）。

解集 不能 包含重复的子集。你可以按 任意顺序 返回解集。
示例 1：

输入：nums = [1,2,3]
输出：[[],[1],[2],[1,2],[3],[1,3],[2,3],[1,2,3]]
示例 2：

输入：nums = [0]
输出：[[],[0]]

链接：https://leetcode.cn/leetbook/read/top-interview-questions-medium/xv67o6/
*/

func subsets(nums []int) [][]int {

	var result = [][]int{}
	result = append(result, []int{})
	traceback(nums, 0, &result)
	return result
}

func traceback(nums []int, index int, result *[][]int) {

	if index >= len(nums) {
		return
	}
	*result = append(*result, []int{nums[index]})
	var next int = 1
	var last = len(*result) - 2
	for next <= last {
		v := (*result)[next]
		v = append(v, nums[index])
		tmp := make([]int, len(v))
		copy(tmp, v)
		*result = append(*result, tmp)
		next++
	}
	traceback(nums, index+1, result)
}

func main() {
	var nums = []int{1, 2, 3} //[]int{1, 2, 3}
	//输出：[[],[1],[2],[1,2],[3],[1,3],[2,3],[1,2,3]]
	fmt.Println(subsets(nums))
}
