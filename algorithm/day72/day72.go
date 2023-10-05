package main

import "fmt"

func main() {

	//fmt.Println(minDistance("horse", "ros"))
	//"intention", word2 = "execution"
	//fmt.Println(minDistance("intention", "execution"))
	//fmt.Println(minDistance("a", "a"))
	// "a"   "ab"
	//fmt.Println(minDistance("a", "ab"))
	fmt.Println(minDistance("abb", "a"))
}

/*
输入：word1 = "horse", word2 = "ros"
输出：3
解释：
horse -> rorse (将 'h' 替换为 'r')
rorse -> rose (删除 'r')
rose -> ros (删除 'e')
*/
// 编辑距离

func min(j, k, l int) int {
	min := j
	if k <= min {
		min = k
	}
	if l <= min {
		min = l
	}
	return min
}
func minDistance(word1 string, word2 string) int {
	m, n := len(word1)+1, len(word2)+1
	if m <= 0 && n <= 0 {
		return 0
	} else if m == 0 {
		return n - 1
	} else if n == 0 {
		return m - 1
	}
	var dp = make([][]int, m)
	for i := 0; i < m; i++ {
		dp[i] = make([]int, n)
	}
	for j := 1; j < n; j++ {
		dp[0][j] = j
	}
	for i := 1; i < m; i++ {
		dp[i][0] = i
	}
	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			if word1[i-1] == word2[j-1] {
				dp[i][j] = dp[i-1][j-1]
			} else {
				// 插入 dp[i][j] = dp[i][j-1] + 1
				// 删除 dp[i][j] = dp[i-1][j] + 1
				// 修改 dp[i][j] = dp[i-1][j-1] + 1
				dp[i][j] = min(dp[i-1][j]+1, dp[i][j-1]+1, dp[i-1][j-1]+1)
			}
		}
	}
	return dp[m-1][n-1]
}
