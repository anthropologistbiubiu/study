package golang_practice

import (
	"container/list"
	"fmt"
)

func containerTest() {

	var ll list.List
	fmt.Println(ll)
	ll.PushBack("a")
	ll.PushBack("b")
	fmt.Println(ll)

}
