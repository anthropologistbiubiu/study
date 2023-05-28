package main

import "fmt"

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
	for x, arr := range board {
		for y, v := range arr {
			if v != word[0] {
				continue
			}
			dfs(aux, board, word, x, y, 0)
			return true
		}
	}
	return false

}

func dfs(aux map[[2]int]struct{}, board [][]byte, word string, x, y, cur int) {

	if x-1 >= 0 && board[x-1][y] == word[cur] {
		if _, ok := aux[[2]int{x, y}]; !ok {
			aux[[2]int{x, y}] = struct{}{}
			dfs(aux, board, word, x, y, cur+1)
		}
		delete(aux, [2]int{x, y})
	}
	if x+1 <= len(board) && board[x+1][y] == word[cur] {
		if _, ok := aux[[2]int{x, y}]; !ok {
			aux[[2]int{x, y}] = struct{}{}
			dfs(aux, board, word, x+1, y, cur+1)
		}
		delete(aux, [2]int{x, y})
	}
	if y-1 >= 0 && board[x][y-1] == word[cur] {
		if _, ok := aux[[2]int{x, y}]; !ok {
			aux[[2]int{x, y}] = struct{}{}
			dfs(aux, board, word, x, y-1, cur+1)
		}
		delete(aux, [2]int{x, y})
	}
	if y+1 < len(board[0]) && board[x][y+1] == word[cur] {
		if _, ok := aux[[2]int{x, y}]; !ok {
			aux[[2]int{x, y}] = struct{}{}
			dfs(aux, board, word, x, y+1, cur+1)
		}
		delete(aux, [2]int{x, y})
	}
	return
}

func main() {
	var test = "abc"
	fmt.Println(string(test[0]))
	fmt.Printf("test[0] %T\n", test[0])
}
