package main

import "fmt"

// 树状数组结构体定义
type BinaryIndexedTree struct {
	tree     []int
	capacity int
}

// Init define
func (bit *BinaryIndexedTree) Init(nums []int) {
	bit.tree = make([]int, len(nums)+1)
	bit.capacity = len(nums)
	for i, v := range nums {
		i += 1
		for i <= bit.capacity {
			bit.tree[i] += v
			i += lowbit(i)
		}
	}

}

func (bit *BinaryIndexedTree) Query(index int) int {

	return 0
}

func (bit *BinaryIndexedTree) QueryRange(start, end int) int {

	return 0
}

func lowbit(x int) int {
	return x & -x
}

func main() {

	nums := []int{1, 3, 5, 7, 9, 10, 22, 5}
	bit := &BinaryIndexedTree{}
	bit.Init(nums)
	fmt.Println(bit.tree)
	fmt.Println(bit.capacity)
}
