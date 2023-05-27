package main

import "fmt"

var ans = make([][]int, 0)

func dfs(cur int, nums []int, curNum []int) {

	//fmt.Println("ans", ans)

	if cur == len(nums) {
		tmp := make([]int, len(curNum))
		copy(tmp, curNum)
		ans = append(ans, tmp)
		return
	}
	for i := cur; i < len(nums); i++ {
		curNum = append(curNum, nums[i])
		dfs(i+1, nums, curNum)
		curNum = curNum[:len(curNum)-1]
		if len(curNum) == 0 {
			continue
		}
		if i == len(nums)-1 {
			tmp := make([]int, len(curNum))
			copy(tmp, curNum)
			ans = append(ans, tmp)
		}
	}
	/*
		tmp := make([]int, len(curNum))
		copy(tmp, curNum)
		ans = append(ans, tmp)
	*/
}

var (
	path []int
	res  [][]int
)

func dfs1(nums []int, start int) {
	tmp := make([]int, len(path))
	copy(tmp, path)
	res = append(res, tmp)

	for i := start; i < len(nums); i++ {
		path = append(path, nums[i])
		dfs1(nums, i+1)
		path = path[:len(path)-1]
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

var testArr = [][]int{}

func test(nums []int) {
	testArr = append(testArr, nums)
	fmt.Println("testArr", testArr)
}

func main() {
	nums := []int{1, 2, 3, 4}
	test(nums)
	//nums[3] = 5
	//fmt.Println(testArr)
	subsets(nums)
	fmt.Println(ans)
}
