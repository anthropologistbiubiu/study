package main

import (
	"fmt"
	"sync"
)

type Set map[string]interface{}

func NewSet() Set {
	return make(Set)
}

func (s Set) Read(reader string) {
	for k, v := range s {
		fmt.Println(reader, k, v)
	}
}

func (s Set) Write(nameArr ...string) {
	for _, name := range nameArr {
		s[name] = struct{}{}
	}
}

func main() {

	set := NewSet()
	cond := sync.NewCond(&sync.Mutex{})

}
