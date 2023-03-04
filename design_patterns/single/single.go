package main

import (
	"fmt"
	"sync"
)

type single struct {
}

var lock sync.Mutex
var singleInstance *single

func NewsingleInstance() *single {
	if singleInstance == nil {
		lock.Lock()
		defer lock.Unlock()
		if singleInstance == nil {
			singleInstance = &single{}
			fmt.Println("Creating single instance now ")
		} else {

			fmt.Println("singleInstance already exists")
		}
	} else {
		fmt.Println("singleInstance already exists")
	}
	return singleInstance
}

func main() {
	for i := 0; i < 10; i++ {
		go NewsingleInstance()
	}
	fmt.Scanln()
}
