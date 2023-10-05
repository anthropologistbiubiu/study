package main

import "fmt"

func main() {

	fmt.Println(minDistance("horse", "ros"))
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
	m, n := len(word1), len(word2)
	if m <= 0 && n <= 0 {
		return 0
	} else if m == 0 {
		return n
	} else if n == 0 {
		return m
	}
	var dp = make([][]int, m)
	for i := 0; i < len(word1); i++ {
		dp[i] = make([]int, n)
	}
	for i := 0; i < len(word1); i++ {
		for j := 0; j < len(word2); j++ {
			if word1[i] == word2[j] {
				dp[i][j] = dp[i-1][j-1] + 1
			} else {
				dp[i][j] = min(dp[i-1][j]+1, dp[i][j-1]+1, dp[i-1][j-1]+1)
			}
		}
	}
	return dp[m-1][n-1]
}
