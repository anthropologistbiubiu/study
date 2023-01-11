package main

import "fmt"

// 买卖股票的最好时机
func maxProfit(prices []int) int {
	// write code here
	var length int = len(prices)
	minPrice := prices[0]
	var totalRevenu int
	for i := 1; i < length; i++ {
		if prices[i] > minPrice {
			totalRevenu += (prices[i] - prices[i-1])
			minPrice += prices[i]
		} else {
			minPrice += prices[i]
		}
	}
	return totalRevenu
}

func main() {
	var arrTest1 []int = []int{1, 83, 74, 26, 63, 37, 25, 63, 28}
	fmt.Println(maxProfit(arrTest1))
	var arrTest2 []int = []int{1, 2, 3, 4, 5}
	fmt.Println(maxProfit(arrTest2))
	var arrTest3 []int = []int{5, 4, 3, 2, 1}
	fmt.Println(maxProfit(arrTest3))
}

//func spiralOrder(matrix [][]int) []int {
//	// write code here
//	var length int = len(matrix)
//	var res []int
//	if length == 0 {
//		return res
//	}
//	var width int = len(matrix[0])
//	var visited [][]int
//	total := length * width
//	for cur := 0; cur < total; {
//		for i := 0; i < length; i++ {
//		}
//		for j := 0; j < width; j++ {
//		}
//		for i := 0; i < length; i++ {
//		}
//		for i := 0; i < length; i++ {
//		}
//	}
//	return res
//}
//
