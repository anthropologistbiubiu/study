package main

import (
	"math"
)

const (
	north = iota
	east
	south
	west
)

// 第一步计算行走的方向 遍历下一个路径  计算下一次的计算 在每个方向上计算最大距离
func robotSim(commands []int, obstacle [][2]int) int {
	mp := make(map[int]int)
	for _, indexs := range obstacle {
		mp[indexs[0]] = indexs[1]
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
					if _, ok := mp[curPos[0]+1]; ok {
						continue
					} else {
						curPos[0] += 1
						if math.Pow(float64(curPos[0]), 2)+math.Pow(float64(curPos[1]), 2) > float64(maxLen) {
							maxLen = int(math.Pow(float64(curPos[0]), 2)) + int(math.Pow(float64(curPos[1]), 2))
						}
					}
				}
			} else if cur == east {
				for i := 1; i < v; i++ {
					_, ok := mp[curPos[0]+1]
					if ok && mp[curPos[0]+1] == curPos[1] {
						continue
					} else {
						curPos[0] += 1
						if math.Pow(float64(curPos[0]), 2)+math.Pow(float64(curPos[1]), 2) > float64(maxLen) {
							maxLen = int(math.Pow(float64(curPos[0]), 2)) + int(math.Pow(float64(curPos[1]), 2))
						}
					}
				}
			} else if cur == south {
				for i := 0; i < v; i++ {
					if _, ok := mp[curPos[0]]; ok {
						if mp[curPos[0]] == curPos[1]-1 {
							continue
						}
					} else {
						curPos[0] -= 1
						if math.Pow(float64(curPos[0]), 2)+math.Pow(float64(curPos[1]), 2) > float64(maxLen) {
							maxLen = int(math.Pow(float64(curPos[0]), 2)) + int(math.Pow(float64(curPos[1]), 2))
						}
					}
				}
			} else if cur == west {
				for i := 0; i < v; i++ {
					if _, ok := mp[curPos[0]-1]; ok {
						if mp[curPos[0]-1] == curPos[1] {
							continue
						}
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
	//var commands = []int{4, -1, 4, -2, 4}
}
