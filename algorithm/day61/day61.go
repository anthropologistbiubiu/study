package main

import "fmt"

var ans = make([][]int, 0)

func dfs(cur int, nums []int, curNum []int) {

	if cur == len(nums) {
		ans = append(ans, curNum)
		fmt.Println("ans", ans)
		return
	}
	for i := cur; i < len(nums); i++ {
		curNum = append(curNum, nums[i])
		dfs(i+1, nums, curNum)
		fmt.Println("cur,i,curNum", cur, i, curNum)
		curNum = curNum[:len(curNum)-1]
		if len(curNum) == 0 {
			continue
		}
		ans = append(ans, curNum)
		fmt.Println("ans", ans)
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
