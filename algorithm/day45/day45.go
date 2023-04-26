package main

import "fmt"

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

func dfs(s []byte, cur int) {
	if cur == len(s)-1 {
		res = append(res, string(s))
	}
	for i := cur; i < len(s); i++ {
		s[i], s[cur] = s[cur], s[i]
		dfs(s, cur+1)
		s[i], s[cur] = s[cur], s[i]
	}
}

var res []string

func main() {
	s := "abc"
	dfs([]byte(s), 0)
	fmt.Println(res)
}
