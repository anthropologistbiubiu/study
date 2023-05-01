package main

import (
	"fmt"
)

//剪绳子

/*
给你一根长度为 n 的绳子，请把绳子剪成整数长度的 m 段（m、n都是整数，n>1并且m>1），每段绳子的长度记为 k[0],k[1]...k[m-1] 。请问 k[0]*k[1]*...*k[m-1] 可能的最大乘积是多少？例如，当绳子的长度是8时，我们把它剪成长度分别为2、3、3的三段，此时得到的最大乘积是18。



输入: 2
输出: 1
解释: 2 = 1 + 1, 1 × 1 = 1



输入: 10
输出: 36
解释: 10 = 3 + 3 + 4, 3 × 3 × 4 = 36
*/

func Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
func cuttingRope(n int) int {
	dp := make([]int, n+1)
	for i := 0; i <= n; i++ {
		if i <= 1 {
			dp[i] = 1
		}
		var maxValue int
		for j := i - 1; j >= 1; j-- {
			if maxValue < Max((i-j)*j, dp[j]*(i-j)) {
				maxValue = Max((i-j)*j, dp[j]*(i-j))
			}
		}
		dp[i] = maxValue
	}
	fmt.Println(dp)
	return dp[n]
}

// (xy) % p = [(x % p)(y % p)] % p 循环求余数的基本原理。
// x^n
func remainder(x, y, p int) int {

	res := int(1)
	for i := 0; i < y; i++ {
		res *= (x % p)
	}
	return res % p
}

/*
	const remainder = (x, a, p = 1000000007) => {
	    let rem = 1;
	    while (a > 0) {
	        if (a % 2) rem = (rem * x) % p;
	        x = (x ** 2) % p
	        a = Math.floor(a / 2);
	    }
	    return rem
	}
*/

//快速幂求余

func remainder1(a, n int) int {

	var res int = 1
	var i int
	for i < n {
		if n&1 == 1 {
			res = res * a % 1000000007
		}
		a = a * a % 1000000007
		n = n >> 1
	}
	return res
}
func cuttingRopePro(n int) int {

	if n <= 3 {
		return n - 1
	}
	m := float64(n % 3)
	i := float64(n / 3)
	var res int
	fmt.Println(m, i)
	switch m {
	case 0:
		res = remainder1(3, int(i)) % 1000000007
	case 1:
		res = remainder1(3, int(i-1)) * 4 % 1000000007 // n = 3m +1 由于 2*2>3*1 故将第 m个3拆成2 + 1  和最后以一个1 拼接为2*2
	case 2:
		res = remainder1(3, int(i)) * 2 % 1000000007
	}
	return res
}
func main() {

	/*
		输入: 2
		输出: 1
		解释: 2 = 1 + 1, 1 × 1 = 1

		输入: 10
		输出: 36
		解释: 10 = 3 + 3 + 4, 3 × 3 × 4 = 36
	*/
	fmt.Println(cuttingRopePro(4))
	fmt.Println(cuttingRopePro(10))
}
