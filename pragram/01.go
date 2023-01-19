package main

import (
	"fmt"
	"sort"
)

// 匿名函数

var Add = func(a, b int) int {
	return a + b
}

func sum(c int, f func(int, int) int) int {
	d := f(c, 3)
	return d
}

func main() {
	re := sum(99, Add)
	fmt.Println(re)
	var sli []int = []int{3, 1, 2}
	var f = func(a, b int) bool {
		return a > b
	}
	sort.Slice(sli, f)
	fmt.Println(sli)
	sort.Slice(sli, func(i, j int) bool {
		return sli[i] < sli[j]
	})
	fmt.Println(sli)
}

// 多个参数和多个返回值
func Swap(a, b int) (int, int) {
	return b, a
}

// 可变数量的参数
// more 对应 []int 切片类型
func Sum(a int, more ...int) int {
	for _, v := range more {
		a += v
	}
	return a
}

/*
Go 语言中的函数可以有多个参数和多个返回值，参数和返回值都是以传值的方式和被调用者交换数据。
在语法上，函数还支持可变数量的参数，可变数量的参数必须是最后出现的参数，可变数量的参数其实是一个切片类型的参数。
*/
