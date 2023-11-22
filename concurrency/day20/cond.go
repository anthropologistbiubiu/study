package main

import (
	"fmt"
	"log"
	"sync"
	"time"
)

type Set map[string]interface{}

var (
	Done = false
)

func NewSet() Set {
	return make(Set)
}

func (s Set) Read(c *sync.Cond, reader string) {
	c.L.Lock()
	defer c.L.Unlock()
	c.Wait()
	for k, v := range s {
		fmt.Println(reader, k, v)
	}
}

func (s Set) Add(name string) {
	s[name] = struct{}{}
}
func (s Set) Has(name string) bool {
	_, ok := s[name]
	return ok
}
func (s Set) Write(c *sync.Cond, nameArr ...string) {
	c.L.Lock()
	log.Print("starting writing")
	for _, name := range nameArr {
		if !s.Has(name) {
			s.Add(name)
		}
	}
	c.L.Unlock()
	c.Broadcast()
}

func main() {

	set := NewSet()
	cond := sync.NewCond(&sync.Mutex{})
	go set.Write(cond, []string{"zhangsan", "lisi"}...)
	go set.Read(cond, "reader1")
	go set.Read(cond, "reader2")
	go set.Read(cond, "reader3")
	time.Sleep(time.Second * 3)
}
