package main

import "fmt"

var ans = make([][]int, 0)

func dfs(cur int, nums []int, curNum []int) {

	if cur == len(nums) {
		ans = append(ans, curNum)
		return
	} else {
		curNum = append(curNum, nums[cur])
		dfs(cur+1, nums, curNum)
		curNum = curNum[:len(curNum)-1]
		ans = append(ans, curNum)
	}
	if cur+1 == len(nums) {
		return
	} else {
		cur++
		dfs(cur+1, nums, curNum)
	}
}

/*
nums = [1,2,3]
输出：[[],[1],[2],[1,2],[3],[1,3],[2,3],[1,2,3]]
*/
func subsets(nums []int) [][]int {

	var curNum = []int{}
	dfs(0, nums, curNum)
	return ans
}

func main() {
	nums := []int{1, 2, 3}
	subsets(nums)
	fmt.Println(ans)
}
