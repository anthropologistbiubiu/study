package main

import (
	"fmt"
)

func Insert(ch byte) {

}

func FirstAppearingOnce() byte {

	return 1
} //

func permutation(s string) []string {
	var res []string
	for i, _ := range s {
		dst := make([]string, len(res))
		copy(dst, res)
		for j, _ := range res {
			res[j] = res[j] + string(s[i])
		}
		if len(res) == 0 {
			res = append(res, string(s[i]))
		} else {
			for j, _ := range dst {
				dst[j] = string(s[i]) + dst[j]
			}
			res = append(res, dst...)
		}
	}
	return res

}

var res []string

func dfs(s []byte, cur int) {
	if cur == len(s)-1 {
		res = append(res, string(s))
	}
	mp := make(map[string]struct{})
	for i := cur; i < len(s); i++ {
		if _, ok := mp[string(s[i])]; ok {
			continue
		} else {
			mp[string(s[i])] = struct{}{}
			s[i], s[cur] = s[cur], s[i]
			dfs(s, cur+1)
			s[i], s[cur] = s[cur], s[i]
		}
	}
}

func main() {
	s := "baa"
	var res []string
	var dfs func(s []byte, cur int)
	dfs = func(s []byte, cur int) {
		if cur == len(s)-1 {
			res = append(res, string(s))
		}
		mp := make(map[string]struct{})
		for i := cur; i < len(s); i++ {
			if _, ok := mp[string(s[i])]; ok {
				continue
			} else {
				mp[string(s[i])] = struct{}{}
				s[i], s[cur] = s[cur], s[i]
				dfs(s, cur+1)
				s[i], s[cur] = s[cur], s[i]
			}
		}

	}
	dfs([]byte(s), 0)
	fmt.Println(res)
}
