package main

func main() {

}

// 编辑距离
func minDistance(word1 string, word2 string) int {

	if len(word1) <= 0 && len(word2) <= 0 {
		return 0
	} else if len(word1) == 0 {
		return len(word2)
	} else if len(word1) == 0 {
		return len(word1)
	}
	return 0
}



