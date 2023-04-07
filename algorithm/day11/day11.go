package main

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
//
//提示：
//你可以假设 k 的值永远是有效的，1 ≤ k ≤ n2
////进阶：
func main() {

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
	//排序
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

func AdjustUp(base []int, start int) {
	j := start       //j来接受pst->size也就是说 j的位置也就是插入节点的起始位置
	i := (j - 1) / 2 //i来当做j的父节点
	for j > 0 {      //开始向上调整，直到j没有父节点{
		if base[j] < base[i] {
			Swap(&(base[i]), &(base[j]))
			j = i           //j追到i的位置作为起始位置，即j以i父节点的位置作为起始调整位置，一步一个脚印向上调整
			i = (j - 1) / 2 //重新定位父节点的位置
		} else {
			break //如果不需要调整，直接结束，开始下一轮的插入调整
		}
	}
}

func MinHeapInsert(php *heap, x int) {
	if php.size < php.capacity {
		php.base[php.size] = x
		AdjustUp(php.base, php.size)
		//调整为小堆结构，从下往上调整
		php.size++ //开始下一轮的插入调整
	}
}
