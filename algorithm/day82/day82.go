package main

func main() {

}

//给你一个整数 n ，请你找出并返回第 n 个 丑数 。

//说明：丑数是只包含质因数 2、3 和/或 5 的正整数；1 是丑数。

func nthUglyNumber(n int) int {

	//
	var dp = make([]int, n)
	dp[0] = 1
	var p2, p3, p5 = dp[0], dp[0], dp[0]
	for i := 1; i < n; i++ {
		dp[i] = min(p2*2, p3*3, p5*5)
	}

	return dp[n-1]
}

func min(a, b, c int) int {
	return 0
}
