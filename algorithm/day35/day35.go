package main

import "fmt"

/*
给你一个由一些多米诺骨牌组成的列表 dominoes。

如果其中某一张多米诺骨牌可以通过旋转 0 度或 180 度得到另一张多米诺骨牌，我们就认为这两张牌是等价的。

形式上，dominoes[i] = [a, b] 和 dominoes[j] = [c, d] 等价的前提是 a==c 且 b==d，或是 a==d 且 b==c。

在 0 <= i < j < dominoes.length 的前提下，找出满足 dominoes[i] 和 dominoes[j] 等价的骨牌对 (i, j) 的数量。

示例：

输入：dominoes = [[1,2],[2,1],[3,4],[5,6]]
输出：1

提示：

1 <= dominoes.length <= 40000
1 <= dominoes[i][j] <= 9
*/
func numEquivDominoPairs1(dominoes [][]int) int {
	var count int
	for i := 0; i < len(dominoes); i++ {
		for j := i + 1; j < len(dominoes); j++ {
			if dominoes[i][0] == dominoes[j][0] && dominoes[i][1] == dominoes[j][1] ||
				dominoes[i][0] == dominoes[j][1] && dominoes[i][1] == dominoes[j][0] {
			}
		}
	}
	return count
}

func numEquivDominoPairs(dominoes [][]int) int {
	mp := make(map[[2]int]int)
	var count int
	for _, v := range dominoes {
		tem1 := [2]int{v[0], v[1]}
		tem2 := [2]int{v[1], v[0]}
		_, ok1 := mp[tem1]
		_, ok2 := mp[tem2]
		if ok1 {
			mp[tem1]++
		} else if ok2 {
			mp[tem2]++
		} else {
			mp[tem1] = 1
		}
	}
	for _, v := range mp {
		if v > 1 {
			count += (v * (v - 1) / 2)
		}
	}
	return count
}

func main() {
	dominoes := [][]int{{1, 2}, {2, 1}, {3, 4}, {5, 6}}
	dominoes1 := [][]int{{1, 2}, {1, 2}, {1, 1}, {1, 2}, {2, 2}}
	fmt.Println(numEquivDominoPairs(dominoes))
	fmt.Println(numEquivDominoPairs(dominoes1))
}
