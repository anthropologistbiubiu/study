package main

import "fmt"

//题目描述
//给定一个排序数组，你需要在 原地 删除重复出现的元素，使得每个元素只出现一次，返回移除后数组的新长度。
//不要使用额外的数组空间，你必须在 原地 修改输入数组 并在使用 O(1) 额外空间的条件下完成。
//示例  1:
//给定数组 nums = [1,1,2],
//函数应该返回新的长度 2, 并且原数组 nums 的前两个元素被修改为 1, 2。
//你不需要考虑数组中超出新长度后面的元素。 示例  2:
//给定 nums = [0,0,1,1,1,2,2,3,3,4],
//函数应该返回新的长度 5, 并且原数组 nums 的前五个元素被修改为 0, 1, 2, 3, 4。
//你不需要考虑数组中超出新长度后面的元素。
//说明:
//为什么返回数值是整数，但输出的答案是数组呢?
//请注意，输入数组是以「引用」方式传递的，这意味着在函数里修改输入数组对于调用者是可见的。
//你可以想象内部操作如下:

func removeDuplicates(arr ...int) []int {
	if len(arr) == 1 {
		return arr
	}
	slow, fast := 0, 0
	for fast < len(arr) {
		if arr[slow] == arr[fast] {
			fast++
		} else {
			for slow+1 != fast {
				arr[fast-1] = arr[fast]
				fast--
			}
			slow++
		}
	}
	return arr[:slow+1]
}
func main() {
	arr := []int{0, 0, 0, 0}
	fmt.Println(removeDuplicates(arr...))
}
