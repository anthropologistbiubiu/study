package day77

/*
给定一个三角形 triangle ，找出自顶向下的最小路径和。

每一步只能移动到下一行中相邻的结点上。相邻的结点
在这里指的是 下标 与 上一层结点下标 相同或者等于 上一层结点下标 + 1 的两个结点。
也就是说，如果正位于当前行的下标 i ，那么下一步可以移动到下一行的下标 i 或 i + 1 。
*/
func minimumTotal(triangle [][]int) int {
	m := len(triangle)
	var dp = make([][]int, m)
	for i := 0; i < m; i++ {
		dp[i] = make([]int, m)
	}
	dp[0][0] = triangle[0][0]
	for i := 1; i < m; i++ {
		dp[i][0] = dp[i-1][0] + triangle[i][0]
	}
	for i := 1; i < m; i++ {
		dp[i][0] = dp[i-1][0] + triangle[i][0]
	}
	return 0
}

// triangle = [[2],[3,4],[6,5,7],[4,1,8,3]]
// 输出：11

func main() {

}
