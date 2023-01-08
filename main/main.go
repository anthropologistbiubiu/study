package main

import (
	"fmt"
	"math"
	"strconv"
)

func Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
func minCost() {
	var n int
	fmt.Scan(&n)
	var arr = make([]int, n)
	var dp = make([]int, n)
	for k := 0; k < n; k++ {
		fmt.Scan(&arr[k])
		dp[k] = 1
	}
	maxlen := dp[0]
	for i := 1; i < n; i++ {
		for j := 0; j < i; j++ {
			if arr[i] > arr[j] {
				dp[i] = Max(dp[j]+1, dp[i])
			}
		}
		maxlen = Max(maxlen, dp[i])
	}
	fmt.Println("maxlen", maxlen)
}

// 最小花费爬楼梯
func minCostClimbingStairs(cost []int) int {
	// write code here
	var length int = len(cost)
	var dp []int = make([]int, length+1)
	if length <= 1 {
		return 0
	}
	dp[0], dp[1] = 0, 0
	for i := 2; i <= length; i++ {
		dp[i] = int(math.Min(float64(dp[i-1]+cost[i-1]), float64(dp[i-2]+cost[i-2])))
	}
	return dp[length]
}

// 打家劫舍
func rob(nums []int) int {
	// write code here
	var length int = len(nums)
	var dp []int = make([]int, length)
	dp[0] = nums[0]
	for i := 1; i < length; i++ {
		if i == 1 {
			dp[i] = int(math.Max(float64(dp[i-1]), float64(nums[i])))
		} else {
			dp[i] = int(math.Max(float64(dp[i-1]), float64(dp[i-2]+nums[i])))
		}
	}
	return dp[length-1]
}

// 把数字翻译成字符串
func stingTonumber(a, b string) int {
	c := a + b
	d, _ := strconv.Atoi(c)
	return d

}

// 力扣版本
func stringTranslate(nums string) int {
	// write code here
	var length int = len(nums)
	var dp []int = make([]int, length)
	dp[0] = 1
	for i := 1; i < length; i++ {
		if i == 1 {
			cur := stingtonumber(string(nums[i-1]), string(nums[i]))
			if cur < 26 && cur >= 10 {
				dp[i] = dp[i-1] + 1
			} else {
				dp[i] = dp[i-1]
			}
		} else {
			cur := stingtonumber(string(nums[i-1]), string(nums[i]))
			if cur <= 26 && cur >= 10 {
				dp[i] = dp[i-1] + dp[i-2]
			} else {
				dp[i] = dp[i-1]
			}
		}
	}
	return dp[length-1]
}
func stingtonumber(a, b string) int {
	c := a + b
	d, _ := strconv.Atoi(c)
	return d

}

// 把数字翻译为字符串 牛客网
func solve(nums string) int {
	// write code here
	if string(nums[0]) == "0" {
		return 0
	}
	var length int = len(nums)
	var dp []int = make([]int, length)
	dp[0] = 1
	for i := 1; i < length; i++ {
		if i == 1 {
			cur := stingtonumber(string(nums[i-1]), string(nums[i]))
			if cur <= 26 && cur > 10 {
				dp[i] = dp[i-1] + 1
			} else {
				dp[i] = dp[i-1]
			}
		} else {
			cur := stingtonumber(string(nums[i-1]), string(nums[i]))
			if cur == 0 {
				return 0
			} else if cur == 10 || cur == 20 {
				dp[i] = dp[i-1]
			} else if cur > 26 && cur%10 == 0 {
				return 0
			} else if cur > 10 && cur <= 26 {
				dp[i] = dp[i-1] + dp[i-2]
			} else {
				dp[i] = dp[i-1]
			}
		}
	}
	return dp[length-1]
}

// 打家劫舍2
func rob2(nums []int) int {
	// write code here
	res1 := rob2core(0, nums)
	res2 := rob2core(1, nums)
	var maximum int
	if res1 > res2 {
		maximum = res1
	} else {
		maximum = res2
	}
	return maximum
}

func rob2core(start int, nums []int) int {
	// write code here
	var length int = len(nums)
	if length == 0 {
		return 0
	}
	if length == 1 {
		return nums[0]
	}
	if length == 2 {
		return int(math.Max(float64(nums[0]), float64(nums[1])))
	}
	if length == 3 {
		var maxVal int
		if nums[0] > nums[1] {
			maxVal = nums[0]
		} else {
			maxVal = nums[1]
		}
		if maxVal < nums[2] {
			maxVal = nums[2]
		}
		return maxVal
	}
	var dp []int = make([]int, length)
	var end int
	if start == 0 {
		end = length - 1
	} else if start == 1 {
		end = length
	}
	for i := start; i < end; i++ {
		if i == 0 {
			dp[i] = nums[i]
		} else if i == 1 {
			if start == 1 {
				dp[i] = nums[i]
			} else {
				dp[i] = int(math.Max(float64(dp[i-1]), float64(nums[i])))
			}
		} else {
			if start == 1 {
				if i == 2 {
					dp[i] = int(math.Max(float64(dp[i-1]), float64(nums[i])))
				} else {
					dp[i] = int(math.Max(float64(dp[i-1]), float64(nums[i]+dp[i-2])))
				}
			} else {
				dp[i] = int(math.Max(float64(dp[i-1]), float64(nums[i]+dp[i-2])))

			}
		}
	}
	return dp[end-1]
}

func main() {
	//var cost = []int{10, 15, 20}
	//var cost1 = []int{1, 100, 1, 1, 1, 100, 1, 1, 100, 1}
	//res := minCostClimbingStairs(cost)
	//fmt.Println(res)
	//result := minCostClimbingStairs(cost1)
	//fmt.Println(result)
	//var rooms []int = []int{1, 2, 3, 4}
	//rs := rob(rooms)
	//fmt.Println(rs)
	//var rooms1 = []int{1, 3, 6}
	//fmt.Println(rob(rooms1))
	//var nums string = "31717126241541717"
	//fmt.Println(solve(nums))
	//fmt.Println(stingtonumber("0", "1"))
	//var nums []int = []int{1, 2, 3, 4}
	var nums []int = []int{19, 43, 94, 4, 34, 33, 91, 75, 38, 79}
	fmt.Println(rob2(nums))

}
