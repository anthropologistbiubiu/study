package main

import "fmt"

func main() {
	fmt.Println(numSubmat([][]int{{1, 0, 1}, {1, 1, 0}, {1, 1, 0}}))
}

func numSubmat(mat [][]int) int {
	width := len(mat)
	long := len(mat[0])
	var count int
	var dp = make([][]bool, width)
	for i, _ := range dp {
		dp[i] = make([]bool, long)
	}
	for i, _ := range dp {
		for j, _ := range dp[i] {
			if i == 0 || j == 0 {
				if mat[i][j] == 1 {
					dp[i][j] = true
					count++
				} else {
					dp[i][j] = false
				}
			} else {
				if dp[i-1][j] && dp[i][j-1] {
					dp[i][j] = true
					count++
				} else {
					dp[i][j] = false
				}
			}
		}
	}
	return count
}
