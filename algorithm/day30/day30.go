package main

import "fmt"

/*
给定两个数组，编写一个函数来计算它们的交集。



示例 1：

输入：nums1 = [1,2,2,1], nums2 = [2,2]
输出：[2]
示例 2：

输入：nums1 = [4,9,5], nums2 = [9,4,9,8,4]
输出：[9,4]


说明：

输出结果中的每个元素一定是唯一的。
我们可以不考虑输出结果的顺序。

*/

func intersection(a []int, b []int) []int {
	var length int
	if len(a) > len(b) {
		length = len(a)
	} else {
		length = len(b)
	}
	mp := make(map[int]int, length)
	for _, v := range a {
		mp[v] = 1
	}
	for _, v := range b {
		if _, ok := mp[v]; ok {
			mp[v]++
		}
	}
	var ret []int
	for k, v := range mp {
		if v > 1 {
			ret = append(ret, k)
		}
	}
	return ret
}
func main() {
	nums1 := []int{1, 2, 2, 1}
	nums2 := []int{2, 2}
	nums3 := []int{4, 9, 5}
	nums4 := []int{9, 4, 9, 8, 4}
	fmt.Println(intersection(nums1, nums2))
	fmt.Println(intersection(nums3, nums4))
}
