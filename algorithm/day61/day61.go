package main

import "fmt"

func dfs(cur int, nums []int, curNum []int, ans *[][]int) {

	if cur == len(nums) {
		tmp := make([]int, len(curNum))
		copy(tmp, curNum)
		*ans = append(*ans, tmp)
		return
	}
	for i := cur; i < len(nums); i++ {
		curNum = append(curNum, nums[i])
		dfs(i+1, nums, curNum, ans)
		curNum = curNum[:len(curNum)-1]
		if len(curNum) == 0 {
			continue
		}
		if i == len(nums)-1 {
			tmp := make([]int, len(curNum))
			copy(tmp, curNum)
			*ans = append(*ans, tmp)
		}
	}
	/*
		tmp := make([]int, len(curNum))
		copy(tmp, curNum)
		ans = append(ans, tmp)
	*/
}

func subsets(nums []int) [][]int {
	var curNum = []int{}
	var ans = make([][]int, 0)
	dfs(0, nums, curNum, &ans)
	return ans
}

var testArr = [][]int{}

func test(nums []int) {
	testArr = append(testArr, []int{2, 3})
	fmt.Println("testArr", testArr)
}

func main() {
	nums := []int{0}
	test(nums)
	//nums[3] = 5
	fmt.Println(nums)
	fmt.Println(subsets(nums))
}
