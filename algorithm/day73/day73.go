package main

import "fmt"

func main() {
	height := []int{0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1}
	fmt.Println(trap(height))
}

// height = [0,1,0,2,1,0,1,3,2,1,2,1]

// 上面是由数组 [0,1,0,2,1,0,1,3,2,1,2,1] 表示的高度图，在这种情况下，可以接 6 个单位的雨水（蓝色部分表示雨水）

func trap(height []int) int {

	length := len(height)
	var totalWater int
	var leftMax = make([]int, length)
	var rightMax = make([]int, length)
	leftMax[0] = height[0]
	for i := 1; i < len(height); i++ {
		leftMax[i] = max(leftMax[i-1], height[i])
	}
	rightMax[length-1] = height[length-1]
	for i := length - 2; i >= 0; i-- {
		rightMax[i] = max(rightMax[i+1], height[i])
	}
	for i := 0; i < length; i++ {
		water := min(leftMax[i], rightMax[i]) - height[i]
		totalWater += water
	}
	return totalWater
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
