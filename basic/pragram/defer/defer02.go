package main

import "fmt"

func main() {
	arr := [5]int{1, 2, 3, 4, 5}
	for i := 0; i < 10; i++ {
		ret, ok := get(arr, i)
		fmt.Printf(" arr[%d]  %t is %d\n", i, ok, ret)
	}
}

func get(arr [5]int, index int) (ret int, ok bool) {
	ok = true

	defer func() {
		err := recover()
		if err != nil {
			ret = -1
			ok = false
		}
	}()
	ret = arr[index]
	return
}
