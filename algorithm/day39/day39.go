package main

import "fmt"

/*
给定一个 n × n 的二维矩阵表示一个图像。

将图像顺时针旋转 90 度。

说明：

你必须在原地旋转图像，这意味着你需要直接修改输入的二维矩阵。请不要使用另一个矩阵来旋转图像。

示例 1:

给定 matrix =
[

	[1,2,3],
	[4,5,6],
	[7,8,9]

],

原地旋转输入矩阵，使其变为:
[

	[7,4,1],
	[8,5,2],
	[9,6,3]

]
示例 2:

给定 matrix =
[

	[ 5, 1, 9,11],
	[ 2, 4, 8,10],
	[13, 3, 6, 7],
	[15,14,12,16]

],

原地旋转输入矩阵，使其变为:
[

	[15,13, 2, 5],
	[14, 3, 4, 1],
	[12, 6, 8, 9],
	[16, 7,10,11]

]


事实上有一个更加巧妙的做法，我们可以巧妙地利用对称轴旋转达到我们的目的，
我们先进行一次以对角线为轴的翻转，然后
再进行一次以水平轴心线为轴的翻转即可。
*/

func matrixPrint(matrix [][]int) {
	length := len(matrix)
	for i := 0; i < length; i++ {
		for j := 0; j < length; j++ {
			fmt.Printf("%3d", matrix[i][j])
		}
		fmt.Println()
	}
	fmt.Println("______________")
}

func rotate(matrix [][]int) {
	length := len(matrix)
	n := length
	for i := 0; i < length; i++ {
		for j := 0; j < n-i-1; j++ {
			matrix[i][j], matrix[n-j-1][n-i-1] = matrix[n-j-1][n-i-1], matrix[i][j]
		}
	}
	for i := 0; i < length/2; i++ {
		for j := 0; j < length; j++ {
			matrix[i][j], matrix[n-1-i][j] = matrix[n-1-i][j], matrix[i][j]
		}
	}
	matrixPrint(matrix)
}
func main() {
	var matrix = [][]int{
		{5, 1, 9, 11},
		{2, 4, 8, 10},
		{13, 3, 6, 7},
		{15, 14, 12, 16},
	}
	matrixPrint(matrix)
	rotate(matrix)

}
