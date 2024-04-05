package main

import (
	"fmt"
	"time"
	"unsafe"
)

func main9() {
	i := 101
	go func(i int) {
		for ; i > 0; i++ {
			fmt.Println(i)
		}
	}(i)
	<-time.After(time.Microsecond * 5)
}

func main8() {
	defer func() {
		r := recover()
		fmt.Println("defer 1", r)
	}()
	defer func() {
		r := recover()
		fmt.Println("defer 2", r)
	}()
	panic("Panic to the reconver function in the second defer function")
}

func main7() {
	mp := make(map[string]int)
	mp["key1"] = 1
	mp["key2"] = 2
	count := **(**int)(unsafe.Pointer(&mp))
	fmt.Println(count)
}

/*

type ST struct{}

func (st ST)String()string{
	return fmt.Sprintf("%s",st.String())
}

func main3(){
	var st = ST{}
	fmt.Println(st.String())
}

func Add(a,b int) int{return a + b}

func main4(){
	v := reflect.ValueOf(Add)
	if v.Kind() !=reflect.Func{
		return
	}
	t := v.Type()
	argv := make([]reflect.Value(),t.NumIn())
	for i:= range argv{
		if t.In(i).Kind() != reflect.Int{
			return
		}
		argv[i] = reflect.ValueOf(i)
	}
	result := v.Call(argv)
	if len(result) != 1 || resrult[0].Kind() != reflect.Int{
		return
	}
	fmt.Println(result[0].Int())
}

*/
