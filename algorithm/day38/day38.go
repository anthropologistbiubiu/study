package main

import (
	"fmt"
	"sort"
)

/*
给定一个字符串数组，将字母异位词组合在一起。字母异位词指字母相同，但排列不同的字符串。

示例:

输入: ["eat", "tea", "tan", "ate", "nat", "bat"]
输出:
[

	["ate","eat","tea"],
	["nat","tan"],
	["bat"]

]
说明：

所有输入均为小写字母。
不考虑答案输出的顺序。
*/
func Sort(str string) string {
	b := []byte(str)
	sort.Slice(b, func(i, j int) bool {
		return b[i] < b[j]
	})
	return string(b)
}
func groupAnagrams(str []string) [][]string {
	mp := make(map[string]int, 10)
	ans := make([][]string, 0)
	for _, v := range str {
		p := Sort(v)
		tem := make([]string, 0)
		if _, ok := mp[p]; !ok {
			mp[p] = len(ans)
			tem = append(tem, v)
			ans = append(ans, tem)
		} else {
			ans[mp[p]] = append(ans[mp[p]], v)
		}
	}
	return ans
}
func main() {
	rst := []string{"eat", "tea", "tan", "ate", "nat", "bat"}
	fmt.Println(groupAnagrams(rst))

}
