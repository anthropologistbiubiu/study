package main

import "fmt"

func subsets(nums []int) [][]int {

	return [][]int{}
}

var ans = make([][]int, 0)

var res = make([]int, 0)

func dfs(cur int, nums []int) {

	if cur == len(nums)-1 {
		ans = append(ans, res)
	}
	dfs(cur+1, nums)
}
func main() {
	fmt.Println()

}
