package main

import "fmt"

// 两个数组的交集
func intersect(arr1 []int, arr2 []int) []int {
	i, j := 0, 0
	var arr = make([]int, 0)
	for i < len(arr1) && j < len(arr2) {
		if arr1[i] == arr2[j] {
			arr = append(arr, arr1[i])
			i++
			j++
		} else if arr1[i] > arr2[j] {
			j++
		} else {
			i++
		}
	}
	return arr
}

func main() {
	var arr1 = []int{1, 2, 3, 5, 6, 9}
	var arr2 = []int{1, 1, 4, 6, 8}
	fmt.Println(intersect(arr1, arr2))
}
