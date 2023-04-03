package main

import (
	"fmt"
	"sync"
)

type A struct {
	Name string
}

func (a *A) Reset() {
	a.Name = ""
}

var pool = sync.Pool{
	New: func() interface{} {
		return new(A)
	},
}

// sync.Pool提供两个接口，Get和Put分别用于从缓存池中获取临时对象，和将临时对象放回到缓存池中：
func main() {
	objA := pool.Get().(*A)
	objA.Reset() // 重置一下对象数据，防止脏数据
	defer pool.Put(objA)
	objA.Name = "test123"
	fmt.Println(objA)
}
