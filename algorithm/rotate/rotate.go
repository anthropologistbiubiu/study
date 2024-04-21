package rotate

import "fmt"

func matrixPrint(matrix [][]int) {
	length := len(matrix)
	for i := 0; i < length; i++ {
		for j := 0; j < length; j++ {
			fmt.Printf("%3d", matrix[i][j])
		}
		fmt.Println()
	}
	fmt.Println("************")
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
