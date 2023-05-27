package main

import "fmt"

/*
给定一个 m x n 二维字符网格 board 和一个字符串单词 word 。如果 word 存在于网格中，返回 true ；否则，返回 false 。
单词必须按照字母顺序，通过相邻的单元格内的字母构成，其中“相邻”单元格是那些水平相邻或垂直相邻的单元格。同一个单元格内的字母不允许被重复使用。

作者：LeetCode
链接：https://leetcode.cn/leetbook/read/top-interview-questions-medium/xvkwe2/
来源：力扣（LeetCode）
著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。
*/

type student struct {
	name string
	age  int
}

/*
	func main() {
		i := 1
		str := "old"

		stu := student{name: "ada", age: 1}

		modify(i, str, stu)
		fmt.Println(i, str, stu.age) //1 old 1
	}

	func modify(i int, str string, stu student) {
		i = 5
		str = "new"
		stu.age = 10
	}

	func main() {
		users := make(map[int]string)
		users[1] = "user1"

		fmt.Printf("before modify: user:%v\n", users[1]) // before modify: user:user1
		modify(users)
		fmt.Printf("after modify: user:%v\n", users[1]) // after modify: user:user2
	}

	func modify(u map[int]string) {
		u[1] = "user2"
	}

func main() {
	arr := make([]int, 0)
	//arr := make([]int, 0, 5)
	arr = append(arr, 1, 2, 3)
	fmt.Printf("outer1: %p, %p, len:%d, capacity:%d\n", &arr, &arr[0], len(arr), cap(arr))
	//modify(arr)
	appendSlice(arr)
	fmt.Printf("outer2: %p, %p, len:%d, capacity:%d\n", &arr, &arr[0], len(arr), cap(arr))
	fmt.Println(arr)
}

func appendSlice(arr []int) {
	fmt.Printf("inner1: %p, %p, len:%d, capacity:%d\n", &arr, &arr[0], len(arr), cap(arr))
	//modify(arr)
	arr = append(arr, 1)
	fmt.Printf("inner2: %p, %p, len:%d, capacity:%d\n", &arr, &arr[0], len(arr), cap(arr))
	//modify(arr) //&arr[0]的地址是否相等，取决于初始化slice的时候的capacity是否足够
}

*/

func main() {
	arr := make([]int, 0)
	arr = append(arr, 1, 2, 3)
	fmt.Printf("outer1: %p, %p\n", &arr, &arr[0])
	modify(arr)
	fmt.Println(arr) // 10, 2, 3
}

func modify(arr []int) {
	fmt.Printf("inner1: %p, %p\n", &arr, &arr[0])
	arr[0] = 10
	fmt.Printf("inner2: %p, %p\n", &arr, &arr[0])
}
