package main

import "fmt"

// .
//给定一个整数数组 nums 和一个目标值 target，请你在该数组中找出和为目标值的那 两个 整数，并返回他们的数组下标。

//你可以假设每种输入只会对应一个答案。但是，你不能重复利用这个数组中同样的元素。

//示例:

//给定 nums = [2, 7, 11, 15], target = 9

//因为 nums[0] + nums[1] = 2 + 7 = 9
//所以返回 [0, 1]

// 不过这样时间复杂度为 O(N^2)，空间复杂度为 O(1)，时间复杂度较高，我们要想办法进行优化。
// 这里我们可以增加一个 Map 记录已经遍历过的数字及其对应的索引值。这样当遍历一个新数字的时候就去 Map 里查询 target 与该数的差值 diff 是否已经在前面的数字中出现过。
// 如果出现过，说明 diff + 当前数 = target，我们就找到了一组答案。

// 关键点
// 求和转换为求差
// 借助 Map 结构将数组中每个元素及其索引相互对应
// 以空间换时间，将查找时间从 O(N) 降低到 O(1)
func twoSum(sum int, arr ...int) []int {
	mp := make(map[int]int, len(arr))
	for i, v := range arr {
		mp[v] = i
	}
	res := []int{}
	for k, v := range mp {
		if _, ok := mp[sum-k]; ok {
			res = append(res, v)
			res = append(res, mp[sum-k])
			return res
		}
	}
	return res
}

func main() {
	nums := []int{2, 7, 4, 5, 9}
	res := twoSum(9, nums...)
	fmt.Println(res)
}
