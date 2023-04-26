package main

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

func dfs(s string, cur int) {
	var res []string
	if cur == len(s)-1 {
		res = append(res, s)
	}
	for i := x; i < len(s); i++ {
		dfs(s, x+1)
		tem := []byte(s)
		tem[x], tem[i] = tem[i], tem[x]
	}
}
