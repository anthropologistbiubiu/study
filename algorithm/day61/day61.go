package main

import "fmt"

var ans = make([][]int, 0)

func dfs(cur int, nums []int, curNum []int) {

	if cur == len(nums)-1 {
		ans = append(ans, curNum)
	}
	curNum = append(curNum, nums[cur])
	dfs(cur+1, nums, curNum)
	ans = append(ans, curNum)
	curNum = curNum[:len(curNum)-1]

	dfs(cur, nums, curNum)
}

func subsets(nums []int) [][]int {

	return [][]int{}
}

func main() {
	fmt.Println()

}
