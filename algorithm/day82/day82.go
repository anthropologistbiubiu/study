package main

import "fmt"

func main() {

	var dp = make([]int, 32)
	var b = -4
	for i := 0; i < 32; i++ {
		dp[i] = b >> i & 1
	}
	fmt.Println(dp)
}
