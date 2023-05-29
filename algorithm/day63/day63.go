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

	var flag bool
	for x, arr := range board {
		for y, v := range arr {
			if v != word[0] {
				continue
			}
			aux := make(map[[2]int]struct{}, 0)
			aux[[2]int{x, y}] = struct{}{}
			fmt.Println("start v,x y", string(v), x, y)
			dfs(aux, board, word, x, y, 0, &flag)
		}
	}
	return flag
}

func dfs(aux map[[2]int]struct{}, board [][]byte, word string, x, y, cur int, flag *bool) {
	fmt.Println("x y cur map ", x, y, cur, aux)
	if cur == len(word)-1 {
		*flag = true
		return
	}
	if cur+1 < len(word) && x-1 >= 0 && board[x-1][y] == word[cur+1] {
		if _, ok := aux[[2]int{x - 1, y}]; !ok {
			aux[[2]int{x - 1, y}] = struct{}{}
			dfs(aux, board, word, x-1, y, cur+1, flag)
		}
		delete(aux, [2]int{x - 1, y})
	}
	if cur+1 < len(word) && x+1 < len(board) && board[x+1][y] == word[cur+1] {
		if _, ok := aux[[2]int{x + 1, y}]; !ok {
			aux[[2]int{x + 1, y}] = struct{}{}
			dfs(aux, board, word, x+1, y, cur+1, flag)
		}
		delete(aux, [2]int{x + 1, y})
	}
	if cur+1 < len(word) && y-1 >= 0 && board[x][y-1] == word[cur+1] {
		if _, ok := aux[[2]int{x, y - 1}]; !ok {
			aux[[2]int{x, y - 1}] = struct{}{}
			dfs(aux, board, word, x, y-1, cur+1, flag)
		}
		delete(aux, [2]int{x, y - 1})
	}
	if cur+1 < len(word) && y+1 < len(board[0]) && board[x][y+1] == word[cur+1] {
		if _, ok := aux[[2]int{x, y + 1}]; !ok {
			aux[[2]int{x, y + 1}] = struct{}{}
			dfs(aux, board, word, x, y+1, cur+1, flag)
		}
		delete(aux, [2]int{x, y + 1})
	}
}

func main() {
	/*
		board := [][]byte{

			{'A', 'B', 'C', 'E'},
			{'S', 'F', 'C', 'S'},
			{'A', 'D', 'E', 'E'},
		}
		test := [][]byte{
			{'C', 'A', 'A'},
			{'A', 'A', 'A'},
			{'B', 'C', 'D'},
		}
		var word1 = "ABCCED"
		var word2 = "SEE"
		var word3 = "ABCB"
		var boards = [][]byte{{'a', 'a'}}
		word4 := "aaa"
		fmt.Println(exist(board, word1))
		fmt.Println(exist(board, word2))
		fmt.Println(exist(board, word3))
		fmt.Println(exist(boards, word4))
		fmt.Println(exist(test, "AAB"))
	*/
	var boardss = [][]byte{
		{'a', 'a', 'a', 'a'},
		{'a', 'a', 'a', 'a'},
		{'a', 'a', 'a', 'a'},
	}
	fmt.Println(exist(boardss, "aaaaaaaaaaaaa"))

}

/// 放弃了改日再会
