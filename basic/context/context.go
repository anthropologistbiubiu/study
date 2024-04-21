package main

import "fmt"

var ans []string

/*
// abcd
固定每个字符串，以每个字符串为首的剩余字符串的排列组合。然后深度遍历剩余字符串的排列组合。至到最后一个字符是该字符串的之后一个字符。
将此时的排列组合加入到结果的队列当中。然后就会退栈道前一个字符，前一个字符。开始交换字符，然后对该字符进行深度优先遍历。
*/
func dfs(s []byte, cur int) {

	if cur == len(s)-1 {
		ans = append(ans, string(s))
		return
	}
	mp := make(map[byte]struct{})
	for i := cur; i < len(s); i++ {
		if _, ok := mp[s[i]]; ok {
			continue
		}
		mp[s[i]] = struct{}{}
		s[cur], s[i] = s[i], s[cur]
		dfs(s, cur+1)
		s[cur], s[i] = s[i], s[cur]
	}
}

func main() {

	dfs([]byte("abb"), 0)
	fmt.Println(ans)
}
