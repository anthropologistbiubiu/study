package main

import "fmt"

func partition(s string) [][]string {

	var length = len(s)
	var dp = make([][]bool, length)
	//var ans = [][]string{}
	for i := 0; i < length; i++ {
		dp[i] = make([]bool, length)
	}
	for i := 0; i < length; i++ {
		dp[i][i] = true
	}
	for i := 0; i < length; i++ {
		for j := 0; j <= i; j++ {
			if i == j {
				dp[i][j] = true
				//st := s[i : j+1]

			}
		}
	}
	return [][]string{}
}

func main() {
	fmt.Println("hello world!")
}
