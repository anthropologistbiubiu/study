package main

import (
	"fmt"
	"math"
)

const (
	north = iota
	east
	south
	west
)

// 第一步计算行走的方向 遍历下一个路径  计算下一次的计算 在每个方向上计算最大距离
func checkCoordinat(x, y int, mp map[[2]int]bool) bool {
	var c = [2]int{x, y}
	if _, ok := mp[c]; ok {
		return true
	}
	return false
}

func robotSim1(commands []int, obstacle [][]int) int {
	mp := make(map[[2]int]bool)
	for _, v := range obstacle {
		var c [2]int = [2]int{v[0], v[1]}
		mp[c] = true
	}
	cur := int(0)
	curPos := [2]int{0, 0}
	maxLen := int(0)
	for _, v := range commands {
		switch v {
		case -1:
			cur = (cur + 1) % 4 // -1 是往右转
		case -2:
			cur = (cur + 4 - 1) % 4 // -2 是往左转
		default:
			if cur == north {
				for i := 0; i < v; i++ {
					if checkCoordinat(curPos[0], curPos[1]+1, mp) {
						break
					} else {
						curPos[1] += 1
						if math.Pow(float64(curPos[0]), 2)+math.Pow(float64(curPos[1]), 2) > float64(maxLen) {
							maxLen = int(math.Pow(float64(curPos[0]), 2)) + int(math.Pow(float64(curPos[1]), 2))
						}
					}
				}
			} else if cur == east {
				for i := 0; i < v; i++ {
					if checkCoordinat(curPos[0]+1, curPos[1], mp) {
						break
					} else {
						curPos[0] += 1
						if math.Pow(float64(curPos[0]), 2)+math.Pow(float64(curPos[1]), 2) > float64(maxLen) {
							maxLen = int(math.Pow(float64(curPos[0]), 2)) + int(math.Pow(float64(curPos[1]), 2))
						}
					}
				}
			} else if cur == south {
				for i := 0; i < v; i++ {
					if checkCoordinat(curPos[0], curPos[1]-1, mp) {
						break
					} else {
						curPos[1] -= 1
						if math.Pow(float64(curPos[0]), 2)+math.Pow(float64(curPos[1]), 2) > float64(maxLen) {
							maxLen = int(math.Pow(float64(curPos[0]), 2)) + int(math.Pow(float64(curPos[1]), 2))
						}
					}
				}
			} else if cur == west {
				for i := 0; i < v; i++ {
					if checkCoordinat(curPos[0]-1, curPos[1], mp) {
						break
					} else {
						curPos[0] -= 1
						if math.Pow(float64(curPos[0]), 2)+math.Pow(float64(curPos[1]), 2) > float64(maxLen) {
							maxLen = int(math.Pow(float64(curPos[0]), 2)) + int(math.Pow(float64(curPos[1]), 2))
						}
					}
				}
			}
		}
	}
	return maxLen
}

/*
-2：向左转 90 度
-1：向右转 90 度
1 <= x <= 9：向前移动 x 个单位长度
在网格上有一些格子被视为障碍物。

示例 1：
输入: commands = [4,-1,3], obstacles = []
输出: 25
解释: 机器人将会到达 (3, 4)
示例 2：
输入: commands = [4,-1,4,-2,4], obstacles = [[2,4]]
输出: 65
解释: 机器人在左转走到 (1, 8) 之前将被困在 (1, 4) 处
*/
func main() {
	comm := []int{-2, 8, 3, 7, -1}
	ob := [][]int{{-4, -1}, {1, -1}, {1, 4}, {5, 0}, {4, 5}, {-2, -1}, {2, -5}, {5, 1}, {-3, -1}, {5, -3}}
	var commands = []int{4, -1, 4, -2, 4}
	var obstacle = [][]int{{2, 4}}
	var c = []int{7, -2, -2, 7, 5}
	var y = [][]int{{-3, 2}, {-2, 1}, {0, 1}, {-2, 4}, {-1, 0}, {-2, -3}, {0, -3}, {4, 4}, {-3, 3}, {2, 2}}
	fmt.Println(robotSim1(commands, obstacle))
	fmt.Println(robotSim1(comm, ob)) //324
	fmt.Println(robotSim1(c, y))     //4
}
