package main

func main() {

}

//给你一个整数 n ，请你找出并返回第 n 个 丑数 。

//说明：丑数是只包含质因数 2、3 和/或 5 的正整数；1 是丑数。

func nthUglyNumber(n int) int {

	//
	var dp = make([]int, n)
	dp[0] = 1
	for i := 1; i < n; i++ {
		dp[i] = min(dp[i-1]*2, dp[i-1]*3, dp[i-1]*5)
	}

	return dp[n-1]
}

func min(a, b, c int) int {
	return 0
}
