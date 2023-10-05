package main

func main() {

}

// 编辑距离
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

		}
	}
	return dp[m-1][n-1]
}
