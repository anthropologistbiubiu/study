package main

import "fmt"

func main() {
	fmt.Println(numSubmat([][]int{{1, 0, 1}, {1, 1, 0}, {1, 1, 0}}))
}

func numSubmat(mat [][]int) int {
	width := len(mat)
	long := len(mat[0])
	var dp = make([][]int, width)
	for i, _ := range dp {
		dp[i] = make([]int, long)
	}
	for i, _ := range dp {
		for j, _ := range dp[i] {
			if i == 0 {
				if mat[i][j] == 1 {

				}
			}
		}
	}
	return 0
}
