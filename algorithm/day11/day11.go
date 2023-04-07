package main

import "fmt"

//给定一个 n x n 矩阵，其中每行和每列元素均按升序排序，找到矩阵中第 k 小的元素。
//请注意，它是排序后的第 k 小元素，而不是第 k 个不同的元素。
//
//
//
//示例：
//
//matrix = [
//   [ 1,  5,  9],
//   [10, 11, 13],
//   [12, 13, 15]
//],
//k = 8,
//
//返回 13。
//
func main() {
	heap := &heap{
		base: make([]int, 100),
		size: 0,
	}
	heap.MinHeapInsert(heap, 9)
	heap.MinHeapInsert(heap, 14)
	heap.MinHeapInsert(heap, 5)
	heap.MinHeapInsert(heap, 8)
	heap.MinHeapInsert(heap, 0)
	heap.MinHeapInsert(heap, 2)
	arr := []int{9, 14, 5, 8, 0, 2}
	heap.HeapSort(heap, arr, 6)
	fmt.Println(arr)

}

type heap struct {
	base     []int
	size     int
	capacity int
}

func Swap(a *int, b *int) {
	temp := *a
	*a = *b
	*b = temp
}
func AdjustDown(base []int, start, n int) {
	i := start
	j := 2*i + 1
	for j < n {
		if j+1 < n && base[j] > base[j+1] {
			j++
		}
		if base[i] > base[j] {
			Swap(&(base[i]), &(base[j]))
			i = j
			j = 2*j + 1

		} else {
			break
		}
	}
}
func (this *heap) HeapSort(pst *heap, ar []int, n int) {
	this.size = n
	curpos := n/2 - 1
	for curpos > 0 {
		AdjustDown(pst.base, 0, n)
		curpos--
	}
	end := n - 1
	for end > 0 {
		Swap(&pst.base[0], &pst.base[end])
		AdjustDown(pst.base, 0, end)
		end--
	}
	k := pst.size - 1
	for j := 0; j < pst.size; j++ {
		ar[j] = pst.base[k]
		k--
	}
}

func AdjutUp(base []int, start int) {
	j := start
	i := (j - 1)
	for j > 0 {
		if base[j] < base[i] {
			Swap(&(base[i]), &(base[j]))
			j = i
			i = (j - 1) / 2
		} else {
			break
		}
	}
}

func (this *heap) MinHeapInsert(php *heap, x int) {
	php.base[php.size] = x
	//AdjustUp(php.base, php.size)
	php.size++
}
