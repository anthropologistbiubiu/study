package test

import (
	"fmt"
	"reflect"
	"sort"
	"unsafe"
)

// 2.字符串转成byte数组，会发生内存拷贝吗？
func test1() {
	a := "aaa"
	ssh := *(*reflect.StringHeader)(unsafe.Pointer(&a))
	b := *(*[]byte)(unsafe.Pointer(&ssh))
	fmt.Printf("%v", b)
}

var testMap map[string]string = make(map[string]string)

func sortmap(testMap map[string]string) {

	for k, v := range testMap {

		fmt.Printf("%s %s\n", k, v)
	}
	sliceTest := make([]string, 1)
	for k, _ := range testMap {
		sliceTest = append(sliceTest, k)
	}
	sort.Strings(sliceTest)
	for _, key := range sliceTest {
		if value, ok := testMap[key]; ok {
			fmt.Println(value)
		}
	}
}
func main9() {
	sortmap(testMap)
}
