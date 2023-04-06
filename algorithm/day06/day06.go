package main

import "fmt"

// 合并两个有序数组
/*
给定两个有序整数数组 nums1 和 nums2，将 nums2 合并到 nums1 中，使得 num1 成为一个有序数组。

说明:

初始化 nums1 和 nums2 的元素数量分别为 m 和 n。
你可以假设 nums1 有足够的空间（空间大小大于或等于 m + n）来保存 nums2 中的元素。
示例:

输入:
nums1 = [1,2,3,0,0,0], m = 3
nums2 = [2,5,6],       n = 3

输出: [1,2,2,3,5,6]
*/
func merge(arr1 []int, arr2 []int) []int {
	if len(arr1) == 0 && len(arr2) == 0 {
		return []int{}
	} else if len(arr1) == 0 {
		return arr2
	} else if len(arr2) == 0 {
		return arr1
	}
	res := make([]int, 0)
	ptr1, ptr2 := 0, 0
	for ptr1 < len(arr1) && ptr2 < len(arr2) {
		if arr1[ptr1] <= arr2[ptr2] {
			res = append(res, arr1[ptr1])
			ptr1++
		} else {
			res = append(res, arr2[ptr2])
			fmt.Println("ptr2", ptr2)
			ptr2++
		}
	}
	if ptr1 < len(arr1) {
		res = append(res, arr1[ptr1:]...)
	}
	if ptr2 < len(arr2) {
		res = append(res, arr2[ptr2:]...)
	}
	return res
}
func main() {

	nums1 := []int{1, 2, 3}
	nums2 := []int{2, 5, 6}
	fmt.Println(merge(nums1, nums2))
}
