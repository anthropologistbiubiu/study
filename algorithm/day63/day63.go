package main

/*
给定一个二维网格和一个单词，找出该单词是否存在于网格中。

单词必须按照字母顺序，通过相邻的单元格内的字母构成，其中“相邻”单元格是那些水平相邻或垂直相邻的单元格。
同一个单元格内的字母不允许被重复使用。

示例:

board =
[

	['A','B','C','E'],
	['S','F','C','S'],
	['A','D','E','E']

]

给定 word = "ABCCED", 返回 true
给定 word = "SEE", 返回 true
给定 word = "ABCB", 返回 false
*/

func exist(board [][]byte, word string) bool {

	aux := make(map[[2]int]struct{}, 0)
	for _, arr := range board {
		for y, v := range arr {

		}
	}

	return true
}

func dfs(aux map[[2]int]struct{}, board [][]byte, word string, x, y int) {

}

func main() {

}
