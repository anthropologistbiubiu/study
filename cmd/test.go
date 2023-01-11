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
			minPrice = prices[i]
		} else {
			minPrice = prices[i]
		}
	}
	return totalRevenu
}

func main() {
	var arrTest []int = []int{7, 1, 5, 3, 6, 4}
	fmt.Println(maxProfit(arrTest))
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
