package golang_practice

import (
	"fmt"
	"time"
	"unsafe"
)

func f1() {
	i := 101
	go func(i int) {
		for ; i > 0; i++ {
			fmt.Println(i)
		}
	}(i)
	<-time.After(time.Microsecond * 5)
}

func f2() {
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

func f3() {
	mp := make(map[string]int)
	mp["key1"] = 1
	mp["key2"] = 2
	count := **(**int)(unsafe.Pointer(&mp))
	fmt.Println(count)
}

type ST struct{}

func (st ST) String() string {
	return fmt.Sprintf("%s", st.String())
}

func f4() {
	var st = ST{}
	fmt.Println(st.String())
}

func Add(a, b int) int { return a + b }
