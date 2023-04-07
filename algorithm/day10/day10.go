package main

import "fmt"

//给你一个包含 n 个整数的数组 nums，判断 nums 中是否存在三个元素 a，b，c ，使得 a + b + c = 0 ？请你找出所有满足条件且不重复的三元组。
//示例：
//
//给定数组 nums = [-1, 0, 1, 2, -1, -4]，
//
//满足要求的三元组集合为：
//[
//  [-1, 0, 1],
//  [-1, -1, 2]
//]

func threeSum(nums ...int) [][]int {
	mp := make(map[int]int, len(nums))
	res := make([][]int, 0)
	for index, value := range nums {
		mp[value] = index
	}
	for i := 0; i < len(nums); i++ {
		tem := make([]int, 0)
		for j := i + 1; j < len(nums); j++ {
			last, ok := mp[0-nums[i]-nums[j]]
			if last > j && ok {
				tem = append(tem, nums[i], 0-nums[i]-nums[j])
			}
		}
		if len(tem) > 0 {
			res = append(res, tem)
		}
	}
	return res
}

func main() {

	nums := []int{-1, 0, 1, 2, -1, -4}
	fmt.Println(threeSum(nums...))

}
